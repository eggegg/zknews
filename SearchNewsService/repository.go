package main

import (
	"context"
	"errors"
	"encoding/json"
	"log"

	"github.com/olivere/elastic"	
	
)

var (
	ErrNotFound = errors.New("Entity not found")
)

type Repository interface {
	Close()
	InsertNews(ctx context.Context, n News) error
	SearchNews(ctx context.Context,  query string, skip int, take int) ([]News, error)
}

type elasticsearchRepository struct {
	client *elastic.Client
}

func NewElasticsearchRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &elasticsearchRepository{client}, nil
}

func (r *elasticsearchRepository) Close()  {
}

func (r *elasticsearchRepository) InsertNews(ctx context.Context, n News) error {
	_, err := r.client.Index().
		Index("news").
		Type("news").
		Id(n.ID).
		BodyJson(n).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *elasticsearchRepository) SearchNews(ctx context.Context, query string, skip int, take int) ([]News, error) {
	
	log.Printf("[search rpc] query:%v, skip:%v, take:%v", query, skip, take)
	

	result, err := r.client.Search().
		Index("news").
		Type("news").
		Query(elastic.NewMultiMatchQuery(query, "title", "content")).
		// Query(
		// 	elastic.NewMultiMatchQuery(query, "title", "content").
		// 		Fuzziness("3").
		// 		PrefixLength(1).
		// 		CutoffFrequency(0.0001),
		// ).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	// log.Println(result.Hits.Hits)
	newsList := []News{}
	for _, hit := range result.Hits.Hits {
		var news News
		if err = json.Unmarshal(*hit.Source, &news); err != nil {
			log.Println(err)
		}
		newsList = append(newsList, news)
	}
	return newsList, nil
}
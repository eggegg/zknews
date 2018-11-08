package main

import (
	"context"
	"time"
	"gopkg.in/mgo.v2/bson"	
)

type Service interface {
	PostNews(ctx context.Context, title, content, author, newsType string, tags []string ) (*News, error)
	GetNews(ctx context.Context, id string) (*News, error)
	GetAllNews(ctx context.Context, skip int, take int) ([]News, error)
}

type News struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Author string `bson:"author" json:"author"`
	NewsType string `bson:"news_type" json:"news_type"`
	Tags []string `bson:"tags" json:"tags"`
	CreatedOn  time.Time  `bson:"createdon,omitempty" json:"createdon,omitempty"`	
}

type newsService struct {
	repository Repository
}

func NewNewsService(r Repository) Service {
	return &newsService{r}
}

func (s *newsService) PostNews(ctx context.Context,title, content, author, newsType string, tags []string ) (*News, error) {
	n := &News{
		ID: bson.NewObjectId(),
		Title: title,
		Content: content,
		Author: author,
		NewsType: newsType,
		Tags: tags,
		CreatedOn: time.Now(),
	}
	if err := s.repository.PutNews(ctx, *n); err != nil {
		return nil, err
	}
	return n, nil
}

func (s *newsService) GetNews(ctx context.Context, id string) (*News, error) {
	return s.repository.GetNewsByID(ctx, id)
}

func (s *newsService) GetAllNews(ctx context.Context, skip int, take int) ([]News, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListNews(ctx, skip, take)
}
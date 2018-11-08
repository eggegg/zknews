package main

import (
	"context"
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
)

var (
	ErrNotFound = errors.New("Entity not found")
)

type Repository interface {
	Close()
	PutNews(ctx context.Context, n News) error
	GetNewsByID(ctx context.Context, id string) (*News, error)
	ListNews(ctx context.Context, skip int, take int) ([]News, error)
}

type mongoRepository struct {
	session *mgo.Session
}

func NewMongoRepository(session *mgo.Session) (Repository, error) {
	return &mongoRepository{session}, nil
}

func (r *mongoRepository) Close()  {
	r.session.Close()
}

func (r *mongoRepository) PutNews(ctx context.Context, n News) error {
	session := r.session.Copy()
	defer session.Close()

	collection := session.DB("zknews").C("news")

	err := collection.Insert(n)
	return err
}

func (r *mongoRepository) GetNewsByID(ctx context.Context, id string) (*News, error ){
	session := r.session.Copy()
	defer session.Close()

	collection := session.DB("zknews").C("news")

	var news News
	err := collection.FindId(bson.ObjectIdHex(id)).One(&news)
	return &news, err
}

func (r *mongoRepository) ListNews(ctx context.Context, skip int, take int) ([]News, error) {
	session := r.session.Copy()
	defer session.Close()

	collection := session.DB("zknews").C("news")

	var newsList []News
	iter := collection.Find(nil).Skip(skip).Limit(take).Iter()
	result := News{}
	for iter.Next(&result) {
		newsList = append(newsList, result)
	}
	return newsList, nil
}
package main

import (
	"context"
	"time"
)

type Service interface {
	InsertNews(ctx context.Context, n *News) error
	SearchNews(ctx context.Context, query string, skip int, take int) ([]News, error)
}


type News struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	NewsType string `json:"news_type"`
	Tags []string `json:"tags"`
	CreatedOn  time.Time  `json:"createdon,omitempty"`	
}


type searchService struct {
	repository Repository
}


func NewSearchService(r Repository) Service {
	return &searchService{r}
}

func (s *searchService) InsertNews(ctx context.Context, n *News) error {
	if err := s.repository.InsertNews(ctx, *n); err != nil {
		return err
	}
	return  nil
}

func (s *searchService) SearchNews(ctx context.Context, query string,skip int, take int) ([]News, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.SearchNews(ctx, query, skip, take)
}
package main

import (
	"context"

	"gopkg.in/mgo.v2/bson"
	"github.com/eggegg/zknews/SportNewsService/pb"
	"google.golang.org/grpc"	
)

type Client struct {
	conn *grpc.ClientConn
	service pb.NewsServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewNewsServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close()  {
	c.conn.Close()
}

func  (c *Client) PostNews(ctx context.Context,title, content, author, newsType string, tags []string ) (*News, error) {
	r, err := c.service.PostNews(
		ctx,
		&pb.PostNewsRequest{
			Title: title,
			Content: content,
			Author: author,
			NewsType: newsType,
			Tags: tags,
		},
	)
	if err != nil {
		return nil, err
	}
	return &News{
		ID: bson.ObjectIdHex(r.News.Id),
		Title: r.News.Title,
		Content: r.News.Content,
		Author: r.News.Author,
		NewsType: r.News.NewsType,
		Tags: r.News.Tags,
	}, nil
}

func (c *Client) GetNews(ctx context.Context, id string) (*News, error) {
	r, err := c.service.GetNews(
		ctx,
		&pb.GetNewsRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &News{
		ID: bson.ObjectIdHex(r.News.Id),
		Title: r.News.Title,
		Content: r.News.Content,
		Author: r.News.Author,
		NewsType: r.News.NewsType,
		Tags: r.News.Tags,
	}, nil
}

func (c *Client) GetAllNews(ctx context.Context, skip int, take int) ([]News, error) {
	r, err := c.service.GetAllNews(
		ctx, 
		&pb.GetAllNewsRequest{
			Skip: int32(skip),
			Take: int32(take),
		},
	)
	if err != nil {
		return nil, err
	}
	newsList := []News{}
	for _, a := range r.Allnews {
		newsList = append(newsList, News{
			ID: bson.ObjectIdHex(a.Id),
			Title: a.Title,
			Content: a.Content,
			Author: a.Content,
			NewsType: a.NewsType,
			Tags: a.Tags,
		})
	}
	return newsList, nil
}
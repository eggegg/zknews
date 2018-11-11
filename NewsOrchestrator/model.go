package main

import (
	"time"
	// "gopkg.in/mgo.v2/bson"
)

type News struct {
	ID string `bson:"id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Author string `bson:"author" json:"author"`
	NewsType string `bson:"news_type" json:"news_type"`
	Tags []string `bson:"tags" json:"tags"`
	CreatedOn  time.Time  `bson:"createdon,omitempty" json:"createdon,omitempty"`	
}
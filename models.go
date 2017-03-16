package main

import (
	"gopkg.in/mgo.v2/bson"
)

type course struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Slug    string        `json:"slug" bson:"slug"`
	Videos  []string      `json:"videos" bson:"videos"`
	Subject string        `json:"subject" bson:"subject"`
	Topic   string        `json:"topic" bson:"topic" `
}

type courses []course

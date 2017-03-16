package main

import (
	"gopkg.in/mgo.v2/bson"
)

type course struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Slug    string        `json:"slug" bson:"slug"`
	Videos  []string      `json:"videos" bson:"videos"`
	Subject bson.ObjectId `json:"subject" bson:"subject"`
	Topic   bson.ObjectId `json:"topic" bson:"topic" `
}

type courses []course

type user struct {
	ID              bson.ObjectId   `json:"id" bson:"_id"`
	Name            string          `json:"name" bson:"name"`
	Email           string          `json:"email" bson:"email"`
	GoogleID        string          `json:"googleID" bson:"googleID"`
	Courses         []bson.ObjectId `json:"courses" bson:"courses"`
	LastLecture     bson.ObjectId   `json:"lastLecture" bson:"lastLecture"`
	LastLectureTime int             `json:"lastLectureTime" bson:"lastLectureTime"`
}

type users []user

type subject struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Slug string        `json:"slug" bson:"slug"`
}

type subjects []subject

type topic struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Slug    string        `json:"slug" bson:"slug"`
	Subject bson.ObjectId `json:"subject" bson:"subject"`
}

type topics []topic

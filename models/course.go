package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Course :
type Course struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Slug    string        `json:"slug" bson:"slug"`
	Videos  []string      `json:"videos" bson:"videos"`
	Subject string        `json:"subject" bson:"subject"`
	Topic   string        `json:"topic" bson:"topic" `
}

// Courses :
type Courses []Course

// // AppLoader :
// type AppLoader struct {
// 	dbSession mongoSession `inject:""`
// }

// GetCourses :
func /*(l *AppLoader)*/ GetCourses() Courses {
	courses := Courses{}

	// if err := l.dbSession.DB("test").C("courses").Find("").One(&courses); err != nil {
	// 	return nil
	// }

	return courses
}

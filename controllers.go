package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func index(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

	courses := courses{}

	if err := c.db.C("courses").Find(bson.M{}).All(&courses); err != nil {
		return http.StatusNotFound, err
	}

	mj, _ := json.Marshal(courses)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	fmt.Fprint(w, string(mj))

	return http.StatusOK, nil
}

func _display(w http.ResponseWriter, o io.Reader) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	buf := &bytes.Buffer{}
	buf.ReadFrom(o)
	fmt.Fprint(w, buf.String())
}

func _parseBody(b io.ReadCloser, t interface{}) {
	buf := &bytes.Buffer{}
	buf.ReadFrom(b)
	json.Unmarshal(buf.Bytes(), t)
}

func _list(db *mgo.Database, t interface{}, c string) (io.Reader, int, error) {
	if err := db.C(c).Find(bson.M{}).All(t); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	mj, err := json.Marshal(t)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return bytes.NewReader(mj), http.StatusOK, nil
}

func _create(db *mgo.Database, t interface{}, c string) (io.Reader, int, error) {
	if err := db.C(c).Insert(t); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	mj, _ := json.Marshal(t)
	return bytes.NewReader(mj), http.StatusOK, nil
}

// Subject

func listSubjects(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	o, s, err := _list(c.db, &subjects{}, "subjects")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

// func getSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

func createSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &subject{ID: bson.NewObjectId()}
	_parseBody(r.Body, &t)
	o, s, err := _create(c.db, t, "subjects")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

// func editSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func deleteSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// // Topic

// func listTopics(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func getTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func editTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func createTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func deleteTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// // Course

// func listCourses(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func getCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func editCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func createCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func deleteCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// // User

// func listUsers(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func getUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func editUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func createUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

// func deleteUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

// }

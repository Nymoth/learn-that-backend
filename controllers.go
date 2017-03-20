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

func _delete(db *mgo.Database, id bson.ObjectId, c string) (int, error) {
	if err := db.C(c).RemoveId(id); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func _edit(db *mgo.Database, id bson.ObjectId, t interface{}, c string) (io.Reader, int, error) {
	if err := db.C(c).UpdateId(id, t); err != nil {
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

func editSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &subject{}
	_parseBody(r.Body, &t)
	o, s, err := _edit(c.db, t.ID, t, "subjects")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func deleteSubject(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &subject{}
	_parseBody(r.Body, &t)
	return _delete(c.db, t.ID, "subjects")
}

// Topic

func listTopics(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	o, s, err := _list(c.db, &topics{}, "topics")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func editTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &topic{}
	_parseBody(r.Body, &t)
	o, s, err := _edit(c.db, t.ID, t, "topics")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func createTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &topic{ID: bson.NewObjectId()}
	_parseBody(r.Body, &t)
	o, s, err := _create(c.db, t, "topics")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func deleteTopic(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &topic{}
	_parseBody(r.Body, &t)
	return _delete(c.db, t.ID, "topics")
}

// Course

func listCourses(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	o, s, err := _list(c.db, &courses{}, "courses")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func editCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &course{}
	_parseBody(r.Body, &t)
	o, s, err := _edit(c.db, t.ID, t, "courses")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func createCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &course{ID: bson.NewObjectId()}
	_parseBody(r.Body, &t)
	o, s, err := _create(c.db, t, "courses")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func deleteCourse(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &course{}
	_parseBody(r.Body, &t)
	return _delete(c.db, t.ID, "courses")
}

// User

func listUsers(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	o, s, err := _list(c.db, &users{}, "users")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func editUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &user{}
	_parseBody(r.Body, &t)
	o, s, err := _edit(c.db, t.ID, t, "users")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func createUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &user{ID: bson.NewObjectId()}
	_parseBody(r.Body, &t)
	o, s, err := _create(c.db, t, "users")
	if err != nil {
		return s, err
	}
	_display(w, o)
	return s, nil
}

func deleteUser(c *context, w http.ResponseWriter, r *http.Request) (int, error) {
	t := &user{}
	_parseBody(r.Body, &t)
	return _delete(c.db, t.ID, "users")
}

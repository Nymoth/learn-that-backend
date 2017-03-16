package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"encoding/json"
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

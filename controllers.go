package main

import (
	"fmt"
	"net/http"

	"encoding/json"
)

// func indexController(c *context) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		courses := courses{}

// 		if err := c.db.DB("test").C("courses").Find("").One(&courses); err != nil {
// 			// return http.StatusNotFound, err
// 			return
// 		}

// 		mj, _ := json.Marshal(courses)

// 		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
// 		fmt.Fprint(w, mj)

// 		// return http.StatusOK, nil
// 		return
// 	}
// }

func index(c *context, w http.ResponseWriter, r *http.Request) (int, error) {

	courses := courses{}

	if err := c.db.DB("test").C("courses").Find("").One(&courses); err != nil {
		return http.StatusNotFound, err
	}

	mj, _ := json.Marshal(courses)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	fmt.Fprint(w, mj)

	return http.StatusOK, nil
}

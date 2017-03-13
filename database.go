package main

import mgo "gopkg.in/mgo.v2"

// GetSession :
func GetSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}

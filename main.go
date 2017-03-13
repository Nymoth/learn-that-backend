package main

import (
	"log"
	"net/http"

	"github.com/facebookgo/inject"
)

func main() {

	mongoSession := GetSession()

	var app RootHandler
	err := inject.Populate(mongoSession, &app)
	if err != nil {
		panic(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"gopkg.in/mgo.v2"
)

type config struct {
	Database databaseConfig `json:"database"`
}

type databaseConfig struct {
	Host string `json:"host"`
}

type context struct {
	cfg *config
	db  *mgo.Session
}

type handler struct {
	*context
	C func(*context, http.ResponseWriter, *http.Request) (int, error)
}

func (h handler) handle(w http.ResponseWriter, r *http.Request) {
	status, err := h.C(h.context, w, r)

	if err != nil {
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func main() {

	cfg := getConfig()
	db := getDBSession(&cfg.Database)

	context := &context{
		cfg: cfg,
		db:  db,
	}

	r := web.New()

	r.Get("/", handler{context, index})

	graceful.ListenAndServe(":8080", r)
}

func getConfig() *config {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var cfg *config
	json.Unmarshal(file, &cfg)

	return cfg
}

func getDBSession(c *databaseConfig) *mgo.Session {

	session, err := mgo.Dial("mongodb://" + c.Host)

	if err != nil {
		panic(err)
	}

	return session
}

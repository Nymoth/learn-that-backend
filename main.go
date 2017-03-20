package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/cors"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"fmt"

	"gopkg.in/mgo.v2"
)

type config struct {
	Server   serverConfig   `json:"server"`
	Database databaseConfig `json:"database"`
}

type serverConfig struct {
	Port string `json:"port"`
}

type databaseConfig struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

type context struct {
	cfg *config
	db  *mgo.Database
}

type handler struct {
	*context
	C func(*context, http.ResponseWriter, *http.Request) (int, error)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h.C(h.context, w, r)

	fmt.Println(r.Method, r.RequestURI, status)

	if err != nil {

		fmt.Println(err.Error())

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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"localhost"},
	})
	r.Use(c.Handler)

	r.Get("/api/subjects", handler{context, listSubjects})
	r.Post("/api/subject", handler{context, createSubject})
	r.Delete("/api/subject", handler{context, deleteSubject})
	r.Put("/api/subject", handler{context, editSubject})

	r.Get("/api/topics", handler{context, listTopics})
	r.Post("/api/topic", handler{context, createTopic})
	r.Delete("/api/topic", handler{context, deleteTopic})
	r.Put("/api/topic", handler{context, editTopic})

	fmt.Println("Server UP at port " + cfg.Server.Port)

	graceful.ListenAndServe(":"+cfg.Server.Port, r)
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

func getDBSession(c *databaseConfig) *mgo.Database {

	session, err := mgo.Dial("mongodb://" + c.Host)

	if err != nil {
		panic(err)
	}

	return session.DB(c.Name)
}

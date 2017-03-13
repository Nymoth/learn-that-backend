package controllers

import (
	"fmt"
	"net/http"

	"github.com/Nymoth/learn-that-backend/models"
	"encoding/json"
)

// Index :
func Index(w http.ResponseWriter, r *http.Request) {

	m := models.GetCourses()
	if m == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

  mj, _  json.Marshal(m)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, mj)
}

package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/shortner", shortner).Methods("GET")
	return m
}

func shortner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Url shortner"))
}

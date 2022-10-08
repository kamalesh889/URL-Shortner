package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/shortner"
	"main.go/store"
)

type response struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

const host = "http://localhost:8080/"

func Router() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/shortner", shortn).Methods("GET")
	m.HandleFunc("/create/short/url", createShorturl).Methods("POST")
	return m
}

func shortn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Url shortner"))
}

func createShorturl(w http.ResponseWriter, r *http.Request) {
	var originalurl string

	err := json.NewDecoder(r.Body).Decode(&originalurl)
	if err != nil {
		fmt.Println("Error in parsing original url ", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	shorturl := shortner.GenerateshorlUrl(originalurl)
	store.SaveUrl(shorturl, originalurl)

	resp := response{
		Message: "Short url created successfully",
		Url:     host + shorturl,
	}

	json.NewEncoder(w).Encode(resp)

}

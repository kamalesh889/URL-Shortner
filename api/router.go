package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/shortner"
	"main.go/store"
)

type request struct {
	Originalurl string
}
type response struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

const host = "http://localhost:8080/"

func Router() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/shortner", shortn).Methods("GET")
	m.HandleFunc("/create/short/url", createShorturl).Methods("POST")
	m.HandleFunc("/{shortUrl}", redirecturl).Methods("GET")
	return m
}

func shortn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Url shortner"))
}

// This Handler implements the shortening of the original url
func createShorturl(w http.ResponseWriter, r *http.Request) {
	var req request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error in parsing original url ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shorturl := shortner.GenerateshorlUrl(req.Originalurl)
	store.SaveUrl(shorturl, req.Originalurl)

	resp := response{
		Message: "Short url created successfully",
		Url:     host + shorturl,
	}

	json.NewEncoder(w).Encode(resp)

}

// This Handler implements the redirection to original url from shorturl
func redirecturl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url, found := vars["shortUrl"]
	if !found {
		fmt.Println("Shorturl didnot present in the request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longurl := store.GetUrl(url)

	http.Redirect(w, r, longurl, http.StatusFound)

}

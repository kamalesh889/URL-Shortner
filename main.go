package main

import (
	"fmt"
	"net/http"

	"main.go/api"
	"main.go/store"
)

func main() {
	fmt.Println("Hello from the service")

	// Intialize db connection
	store.IntializeConnection()

	// Router implementation
	mux := api.Router()
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	"main.go/api"
)

func main() {
	fmt.Println("Hello to the service")
	mux := api.Router()
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}

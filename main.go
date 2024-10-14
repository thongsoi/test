package main

import (
	"net/http"
	"your_project/handler" // Import your handler

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fetch-categories", handler.FetchCategoriesHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

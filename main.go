package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/test/database"
	"github.com/thongsoi/test/handler"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	// Route for rendering the form (GET)
	r.HandleFunc("/", handler.FormHandler).Methods("GET")

	// Route for form submission (POST)
	r.HandleFunc("/submit-category", handler.SubmitCategoryHandler).Methods("POST")

	http.ListenAndServe(":9000", r)
}

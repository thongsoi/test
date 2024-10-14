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

	// Route for fetching categories (HTMX)
	r.HandleFunc("/fetch-categories", handler.FetchCategoriesHandler).Methods("GET")

	// Route for form submission
	r.HandleFunc("/submit-category", handler.SubmitCategoryHandler).Methods("POST")

	// Serve the HTML file (e.g., form.html)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":9000", r)
}

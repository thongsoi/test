package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/test/Internal/order"
	"github.com/thongsoi/test/database"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	// Route for rendering the form (GET)
	r.HandleFunc("/", order.FormHandler).Methods("GET")

	// Route for form submission (POST)
	r.HandleFunc("/submit-category", order.SubmitCategoryHandler).Methods("POST")

	http.ListenAndServe(":9000", r)
}

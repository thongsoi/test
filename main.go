package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/test/database" // Import your db package
	"github.com/thongsoi/test/handler"  // Import your handler package
)

func main() {
	// Initialize the database connection
	database.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Define the route for fetching categories
	r.HandleFunc("/fetch-categories", handler.FetchCategoriesHandler).Methods("GET")

	// Serve the HTML file (e.g., form.html) at the root URL
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Start the server on port 8080
	http.ListenAndServe(":9000", r)
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thongsoi/test/database"
	"github.com/thongsoi/test/internal/order"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	// Route for form submission (POST)
	r.HandleFunc("/", order.FormHandler).Methods("GET")
	r.HandleFunc("/order", order.SubmitOrderHandler).Methods("POST")

	http.ListenAndServe(":9000", r)
}

package test

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route for fetching options
	r.HandleFunc("/options", handler.OptionsHandler)

	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}

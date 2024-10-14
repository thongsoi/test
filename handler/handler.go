package test

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/thongsoi/biomassx/database"
)

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the data from PostgreSQL
	db := database.InitDB() // Replace with your DB instance
	options, err := FetchOptions(db)
	if err != nil {
		http.Error(w, "Unable to fetch options", http.StatusInternalServerError)
		return
	}

	// Use a simple template to render the options
	tmpl, err := template.ParseFiles(path.Join("templates", "options.html"))
	if err != nil {
		log.Fatal(err)
	}

	// Serve the template with options data
	tmpl.Execute(w, options)
}

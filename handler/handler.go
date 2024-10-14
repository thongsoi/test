package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/thongsoi/test/database"
	"github.com/thongsoi/test/repository" // Import your repository package
	// Import your database connection
)

func FetchCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Get categories from the database
	categories, err := repository.FetchCategories(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch categories", http.StatusInternalServerError)
		return
	}

	// Parse the template for the <option> elements
	tmpl, err := template.New("categories").Parse(`
		{{range .}}
			<option value="{{.ID}}">{{.Name}}</option>
		{{end}}
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with the fetched categories
	err = tmpl.Execute(w, categories)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

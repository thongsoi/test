package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/thongsoi/test/database"
	"github.com/thongsoi/test/repository"
)

func FetchCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := repository.FetchCategories(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch categories", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("categories").Parse(`
        {{range .}}
            <option value="{{.ID}}">{{.EnName}}</option>
        {{end}}
    `)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, categories)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func SubmitCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the selected category ID
	categoryID := r.FormValue("category")

	// Convert category ID to an integer (optional)
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Do something with the selected category, e.g., save it or display it
	fmt.Fprintf(w, "Selected Category ID: %d", categoryIDInt)
}

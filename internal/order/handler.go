package order

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/thongsoi/test/database"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	markets, err := FetchMarkets(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch markets", http.StatusInternalServerError)
		return
	}

	// Parse the Go template
	tmpl, err := template.ParseFiles("templates/order.html") // Adjust path if necessary
	if err != nil {
		log.Fatal(err)
	}

	// Pass categories to the template
	data := struct {
		Markets []Market
	}{
		Markets: markets,
	}

	// Execute the template and render the HTML page
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func SubmitMarketIDHandler(w http.ResponseWriter, r *http.Request) {
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

	// Retrieve the selected category from the form data (by name attribute "category")
	marketID := r.FormValue("marketID")

	// Convert category ID to an integer (optional)
	marketIDInt, err := strconv.Atoi(marketID)
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	// Process the selected category (for now, we will just display it in the response)
	log.Printf("Selected market ID: %d", marketIDInt)

	// Respond with a confirmation message (or redirect to another page, etc.)
	fmt.Fprintf(w, "You selected market with ID: %d", marketIDInt)
}

/*
func FormHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := FetchCategories(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch categories", http.StatusInternalServerError)
		return
	}

	// Parse the Go template
	tmpl, err := template.ParseFiles("templates/form.html") // Adjust path if necessary
	if err != nil {
		log.Fatal(err)
	}

	// Pass categories to the template
	data := struct {
		Categories []Category
	}{
		Categories: categories,
	}

	// Execute the template and render the HTML page
	err = tmpl.Execute(w, data)
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

	// Retrieve the selected category from the form data (by name attribute "category")
	categoryID := r.FormValue("category")

	// Convert category ID to an integer (optional)
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Process the selected category (for now, we will just display it in the response)
	log.Printf("Selected Category ID: %d", categoryIDInt)

	// Respond with a confirmation message (or redirect to another page, etc.)
	fmt.Fprintf(w, "You selected category with ID: %d", categoryIDInt)
}
*/

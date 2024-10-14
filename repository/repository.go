package repository

import (
	"database/sql"
	"log"
)

// Assuming you have a Category struct
type Category struct {
	ID   int
	Name string
}

func FetchCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		categories = append(categories, category)
	}
	return categories, nil
}

package order

import (
	"database/sql"
	"log"
)

func FetchCategories(db *sql.DB) ([]Category, error) {
	// Order by category name in ascending order
	rows, err := db.Query("SELECT id, en_name FROM categories ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.EnName)
		if err != nil {
			log.Println(err)
			continue
		}
		categories = append(categories, category)
	}
	return categories, nil
}

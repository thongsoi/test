package repository

import (
	"database/sql"
	"log"

	"github.com/thongsoi/test/models"
)

func FetchCategories(db *sql.DB) ([]models.Category, error) {
	// Order by category name in ascending order
	rows, err := db.Query("SELECT id, en_name FROM categories ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.EnName)
		if err != nil {
			log.Println(err)
			continue
		}
		categories = append(categories, category)
	}
	return categories, nil
}

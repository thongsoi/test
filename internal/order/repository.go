package order

import (
	"database/sql"
	"log"
)

func FetchMarkets(db *sql.DB) ([]Market, error) {
	// Order by market name in ascending order
	rows, err := db.Query("SELECT id, en_name FROM markets ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []Market
	for rows.Next() {
		var market Market
		err := rows.Scan(&market.ID, &market.EnName)
		if err != nil {
			log.Println(err)
			continue
		}
		markets = append(markets, market)
	}
	return markets, nil
}

func FetchSubmarkets(db *sql.DB) ([]Submarket, error) {
	// Order by submaket name in ascending order
	rows, err := db.Query("SELECT id, en_name FROM submarkets ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submarkets []Submarket
	for rows.Next() {
		var submaket Submarket
		err := rows.Scan(&submaket.ID, &submaket.EnName)
		if err != nil {
			log.Println(err)
			continue
		}
		submarkets = append(submarkets, submaket)
	}
	return submarkets, nil
}

/*
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
*/

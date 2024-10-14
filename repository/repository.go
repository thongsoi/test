package test

import (
	"database/sql"
	"log"
)

// Fetch options from PostgreSQL
func FetchOptions(db *sql.DB) ([]Option, error) {
	var options []Option

	query := `SELECT id, name FROM options_table`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var option Option
		if err := rows.Scan(&option.ID, &option.Name); err != nil {
			log.Fatal(err)
		}
		options = append(options, option)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return options, nil
}

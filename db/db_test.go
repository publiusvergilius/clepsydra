package db

import (
	"database/sql"
	"log"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

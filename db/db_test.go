package db

import (
	"database/sql"
	"log"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:") // ":memory:" or "./data.db"
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SetupDevDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:") // ":memory:" or "./data.db"
	if err != nil {
		log.Fatal(err)
	}

	return db
}

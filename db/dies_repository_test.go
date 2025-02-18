package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// diretamente, o database é reiniciado a cada teste

var db = setupTestDB()

func TestDiesRepository_Initialization(t *testing.T) {
	var name string
	_, err := NewDiesRepository(db)

	// dies, err := repo.GetAll()

	if err != nil {
		t.Fatalf("Expected no error while creating repo, got %v", err)
	}

	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='dies'").Scan(&name)

	if err == sql.ErrNoRows {
		t.Fatalf("Expected table %s to be created.", name)
	}

	if err != nil {
		t.Fatalf("Expected no error while verifying columns in the table.")
	}
}

func TestDiesRepository_GetAll_Empty(t *testing.T) {}

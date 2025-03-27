package db

import (
	"testing"
	"time"
)

var testDB = SetupDevDB()

var dr = DiesRepository{}

func TestCreateDies(t *testing.T) {
	t.Run("create from time.Now", func(t *testing.T) {
		dies := Dies{}
		dies.SetDate(time.Now())
		dies.id = 1

		dr.Create(testDB)
		dr.Save(testDB, dies)

		assertExists(t, dies)

		dies2 := Dies{}
		dies2.SetDate(time.Now())
		err := dr.Save(testDB, dies2)

		if err == nil {
			t.Errorf("was told to error on saving repeated `dies`")
		}
	})

	// Save user defined date
	t.Run("create user defined date from string", func(t *testing.T) {

		dateStr := "2025-03-28"
		err := dr.Save(testDB, Dies{id: 2, date: dateStr})
		assertNotError(t, err)

		got, err := dr.FindByDate(testDB, dateStr)
		assertNotError(t, err)

		want := Dies{id: 2, date: dateStr}

		if want != got {
			t.Fatalf("want %q, got %q", want, got)
		}
	})
}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("was not told to error: %q", err)
	}
}

func assertExists(t *testing.T, want Dies) {

	got, err := dr.FindById(testDB, want.id)
	if err != nil {
		t.Errorf("was not told to error: %q", err)
	}

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

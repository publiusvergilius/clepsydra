package db

import (
	"time"
)

// represents a row of a day that contains it's tasks
// each tasks has a relation to it's day
// many-to-one
type Dies struct {
	id   uint   `sqlite3:"id"`
	date string `sqlite3:"date"` // 31-01-2025
}

func (d *Dies) GetID() uint {
	return d.id
}

func (d *Dies) GetDate() string {
	return d.date
}

func (d *Dies) SetDate(date time.Time) {
	d.date = date.UTC().Format(time.DateOnly)
}

type DiesRepository struct{}

func (d *DiesRepository) Create(db DB) (Result, error) {
	sqlStmt := `create table "dies" (id integer not null unique primary key autoincrement, date text unique);`

	return db.Exec(sqlStmt)
}

func (d *DiesRepository) FindAll(db DB) ([]Dies, error) {

	diesList := []Dies{}
	querySQL := "select id, date from dies"
	rows, err := db.Query(querySQL)
	if err != nil {
		return diesList, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var date string

		err := rows.Scan(&id, &date)
		if err != nil {
			return []Dies{}, err
		}

		newDies := Dies{id: id}
		parsedDate, err := time.Parse(time.DateOnly, date)

		if err != nil {
			return []Dies{}, err
		}

		newDies.SetDate(parsedDate)
		diesList = append(diesList, newDies)
	}

	return diesList, nil
}

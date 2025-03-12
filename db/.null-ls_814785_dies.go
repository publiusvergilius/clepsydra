package db

import (
	"time"
)

// represents a row of a day that contains it's tasks
// each tasks has a relation to it's day
// many-to-one
type Dies struct {
	id      uint   `sqlte3:"id"`
	date    string `sqlte3:"date"` // 31-01-2025
}

func (d *Dies) GetID() uint {
	return d.id
}


func (d *Dies) GetDate() string {
	return d.date
}

func (d *Dies) SetDate(date time.Time) {
	d.date = date.Format(time.DateOnly)
}

type DiesRepository struct{}

func NewDiesRepository(db DB) (Result, error) {
	sqlStmt := `create table dies (id integer not null primary key autoincrement, titulum text, date text unique);`

	return db.Exec(sqlStmt)
}

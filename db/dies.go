package db

import (
	"encoding/json"
	"time"
)

// represents a row of a day that contains it's tasks
// each tasks has a relation to it's day
// many-to-one
type Dies struct {
	id   uint   `sqlite3:"id"`
	date string `sqlite3:"date"` // 31-01-2025
}

func (d Dies) GetID() uint {
	return d.id
}

func (d Dies) GetDate() string {
	return d.date
}

func (d *Dies) SetDate(date time.Time) {
	d.date = date.UTC().Format(time.DateOnly)
}

func (d Dies) ToString() (string, error) {
	type JsonStruct struct {
		Id   uint   `json:"id"`
		Date string `json:"date"`
	}
	newDies := JsonStruct{
		Id:   d.GetID(),
		Date: d.GetDate(),
	}

	jsonStr, err := json.Marshal(newDies)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

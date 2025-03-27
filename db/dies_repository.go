package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

/** !TODO: create a CreateDate method: (time.Time) -> string
 */

type DiesRepository struct{}

func (DiesRepository) Create(db DB) (Result, error) {
	sqlStmt := `create table "dies" 
		    (id integer not null primary key autoincrement, 
		    date text unique);`

	return db.Exec(sqlStmt)
}

func (DiesRepository) FindAll(db DB) ([]Dies, error) {

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

func (DiesRepository) FindById(db DB, id uint) (Dies, error) {
	var dies Dies

	stmt := `select id, date from dies where id = ?`
	err := db.QueryRow(stmt, id).Scan(&dies.id, &dies.date)

	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := fmt.Sprintf("no dies found with ID %d", id)
			return Dies{}, errors.New(errMessage)
		}
		return Dies{}, err
	}

	return dies, nil
}

func (DiesRepository) Save(db DB, d Dies) error {
	_, err := db.Exec(`insert into dies(date) values (?)`,
		d.GetDate())

	if err != nil {
		return err
	}

	return nil
}

func (DiesRepository) Delete(db DB, id uint) (int64, error) {
	/*
		result, err := db.Exec("delete from quartum where id = ?;", id)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		rowsAffected, err := result.RowsAffected()

		if err != nil {
			return rowsAffected, err
		}
		return rowsAffected, nil
	*/
	return 0, nil
}

func (DiesRepository) FindQuarta(db DB, id uint) ([]Quartum, error) {
	var list []Quartum

	stmt := `
		 select id, titulum, hora, pars, dies_id
		 from quartum
		 where dies_id = ?;
		`

	rows, err := db.Query(stmt, id)
	if err != nil {
		log.Default().Fatalln("error on FindAllQuarta: ", err)
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		q := Quartum{}

		err := rows.Scan(&q.id, &q.titulum, &q.hora, &q.pars, &q.dies_id)
		if err != nil {
			return []Quartum{}, err
		}

		list = append(list, q)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return list, nil
}

func (DiesRepository) FindByDate(db DB, dateStr string) (Dies, error) {
	stmt := `select id, date from dies where date = ?`
	dies := Dies{}

	date, err := MakeDateFromString(dateStr)
	if err != nil {
		return dies, err
	}

	rows, err := db.Query(stmt, date.Format(time.DateOnly))
	if err != nil {
		return dies, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dies.id, &dies.date)
		if err != nil {
			return Dies{}, err
		}
	}

	return dies, nil
}

func MakeDateFromString(dateStr string) (time.Time, error) {
	layout := "2006-01-02"

	date2, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date2, nil
}

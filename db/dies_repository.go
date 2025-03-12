package db

import "time"

type DiesRepository struct{}

func (DiesRepository) Create(db DB) (Result, error) {
	sqlStmt := `create table "dies" (id integer not null unique primary key autoincrement, date text unique);`

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

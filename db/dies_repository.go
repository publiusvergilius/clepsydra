package db

import (
	"log"
	"time"
)

type DiesRepository struct{}

func (DiesRepository) Create(db DB) (Result, error) {
	sqlStmt := `
		    create table "dies" 
		    (id integer not null unique primary key autoincrement, 
		    date text unique);
		    `

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
	/*
		var quartum Quartum

		stmt := `select id, titulum, hora, pars, dies_id from quartum where id = ?`
		err := db.QueryRow(stmt, id).Scan(&quartum.id, &quartum.titulum, &quartum.hora, &quartum.pars, &quartum.dies_id)

		if err != nil {
			if err == sql.ErrNoRows {
				errMessage := fmt.Sprintf("no quartum found with ID %d", id)
				return Quartum{}, errors.New(errMessage)
			}
			return Quartum{}, err
		}

		return quartum, nil
	*/
	return Dies{}, nil
}

func (DiesRepository) Save(db DB, q Dies) error {

	/*
		db.Exec(`insert into quartum(
				pars, titulum, hora, dies_id)
				values (?, ?, ?, ?)`,
			q.GetPars(), q.GetTitulum(), q.GetHora(), q.GetDiesId())
		return nil
	*/

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

func (DiesRepository) FindAllQuarta(db DB, id uint) ([]Quartum, error) {
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

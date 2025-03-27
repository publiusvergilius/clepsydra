package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type JsonQuartum struct {
	Titulum string `json:"titulum"`
	Pars    uint   `json:"pars"`
	Hora    string `json:"hora"`
	Dies_id uint   `json:"dies_id"`
}

type QuartumRepository struct{}

func (QuartumRepository) Create(db DB) (Result, error) {
	/** Child table e Foreign key constraint*/
	sqlStmt := `create table "quartum" (
		id integer not null primary key autoincrement, 
		titulum varchar(50) not null, 
		pars tinyint not null, 
		hora text, 
		dies_id integer not null,
		unique(titulum, pars, dies_id),
		foreign key (dies_id) references dies(id) on delete cascade
	);`

	return db.Exec(sqlStmt)
}

func (QuartumRepository) FindAll(db DB) ([]Quartum, error) {
	quartumList := []Quartum{}
	querySQL := "select id, titulum, hora, pars, dies_id from quartum"
	rows, err := db.Query(querySQL)
	if err != nil {
		return quartumList, err
	}
	defer rows.Close()

	for rows.Next() {
		var newQuartum Quartum

		err := rows.Scan(&newQuartum.id, &newQuartum.titulum, &newQuartum.hora, &newQuartum.pars, &newQuartum.dies_id)

		if err != nil {
			if err == sql.ErrNoRows {
				errMessage := fmt.Sprintf("no quartum found: %q", err)
				return []Quartum{}, errors.New(errMessage)
			}
			return []Quartum{}, err
		}

		quartumList = append(quartumList, newQuartum)
	}

	return quartumList, nil
}

func (QuartumRepository) FindById(db DB, id uint) (Quartum, error) {
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
}

func (QuartumRepository) Save(db DB, q Quartum) error {

	db.Exec(`insert into quartum(
			pars, titulum, hora, dies_id) 
			values (?, ?, ?, ?)`,
		q.GetPars(), q.GetTitulum(), q.GetHora(), q.GetDiesId())
	return nil
}

func (QuartumRepository) Delete(db DB, id uint) (int64, error) {

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
}

func (QuartumRepository) FindByDies(db DB, id uint) ([]Quartum, error) {
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

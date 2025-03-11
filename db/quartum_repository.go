package db

import (
	"database/sql"
	"errors"
	"fmt"
)

type QuartumRepository struct{}

func (QuartumRepository) Create(db DB) (Result, error) {
	/** Child table e Foreign key constraint*/
	sqlStmt := `create table "quartum" (
		id integer not null unique primary key autoincrement, 
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

		fmt.Println(newQuartum)
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

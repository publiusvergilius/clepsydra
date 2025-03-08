package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Quartum struct {
	id      uint
	titulum string
	pars    uint8
	hora    string
	dies_id uint
}

func (q *Quartum) GetTitulum() string {
	return q.titulum
}

func (q *Quartum) SetTitulum(titulum string) error {

	if len(titulum) > 60 {
		return errors.New("titulum não pode ser maior que 60 caracteres.")
	}
	q.titulum = titulum
	return nil
}

func (q *Quartum) GetHora() string {
	return q.hora
}
func (q *Quartum) SetHora(hora time.Time) {
	q.hora = hora.UTC().Format(time.TimeOnly)
}

func (q *Quartum) GetPars() uint8 {
	return q.pars
}

func (q *Quartum) SetPars(pars uint8) {
	q.pars = pars
}

func (q *Quartum) GetDiesId() uint {
	return q.dies_id
}

func (q *Quartum) SetDiesId(id uint) {
	q.dies_id = id
}

type QuartumRepository struct{}

func (q *QuartumRepository) Create(db DB) (Result, error) {
	sqlStmt := `create table "quartum" (
		id integer not null unique primary key autoincrement, 
		titulum varchar(50) not null, 
		pars tinyint not null, 
		hora text, 
		dies_id integer not null
	);`

	return db.Exec(sqlStmt)
}

func (q *QuartumRepository) FindAll(db DB) ([]Quartum, error) {
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

func (q *QuartumRepository) FindById(db DB, id uint) (Quartum, error) {
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

func (q *QuartumRepository) Save(db DB, quartum Quartum) error {

	db.Exec(`insert into quartum(
			pars, titulum, hora, dies_id) 
			values (?, ?, ?, ?)`,
		quartum.GetPars(), quartum.GetTitulum(), quartum.GetHora(), quartum.GetDiesId())
	return nil
}

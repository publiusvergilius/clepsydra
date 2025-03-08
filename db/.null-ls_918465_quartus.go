package db

import (
	"errors"
	"time"
)

type Quartum struct {
	id      uint
	titulum string
	pars    uint8
	hora    string
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

type QuartumRepository struct{}

func (q *QuartumRepository) New(db DB) (Result, error) {
	sqlStmt := `create table quartum (id integer not null unique primary key autoincrement, titulum varchar(50) not null, pars tinyint not null, hora text);`

	return db.Exec(sqlStmt)
}

func (q *QuartumRepository) FindAll(db DB) ([]Quartum, error) {
	quartumList := []Quartum{}
		querySQL := "select id, titulum, hora, pars from quartum"
		rows, err := db.Query(querySQL)
		if err != nil {
			return quartumList, err
		}
		defer rows.Close()

		for rows.Next() {
			var id uint
			var titulum string
			var hora string
			var pars uint8

			err := rows.Scan(&id, &titulum, &hora, &pars)

			if err != nil {
				return quartumList, err
			}

			newQuartum := Quartum{id: id, hora: hora}
			newQuartum.SetPars(pars)
			newQuartum.SetTitulum(titulum)
			// newQuartum.SetHora(hora)

			quartumList = append(quartumList, newQuartum)
		}
	return quartumList, nil
}

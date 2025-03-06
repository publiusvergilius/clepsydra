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

func (q *QuartumRepository) newQuartumRepository(db DB) (Result, error) {
	sqlStmt := `create table quartum (id integer not null unique primary key autoincrement, titulum varchar(50) not null, pars tinyint not null, hora text);`

	return db.Exec(sqlStmt)
}

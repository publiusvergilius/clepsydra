package db

import (
	"encoding/json"
	"errors"
	"time"
)

type Quartum struct {
	id      uint
	titulum string
	pars    uint8
	hora    string
	prazo   string
	dies_id uint
}

func (q Quartum) GetID() uint {
	return q.id
}

func (q *Quartum) GetTitulum() string {
	return q.titulum
}

func (q *Quartum) SetTitulum(titulum string) error {

	if len(titulum) > 60 {
		return errors.New("titulum não pode ser maior que 60 caracteres")
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

func (q *Quartum) GetPrazo() string {
	return q.prazo
}
func (q *Quartum) SetPrazo(hora time.Time) {
	q.prazo = hora.UTC().Format(time.TimeOnly)
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

func (q *Quartum) ToString() (string, error) {
	type JsonQuartum struct {
		Id      uint   `json:"id"`
		Titulum string `json:"titulum"`
		Pars    uint8  `json:"pars"`
		Hora    string `json:"hora"`
		Dies_id uint   `json:"dies_id"`
	}
	newQuartum := JsonQuartum{
		Id:      q.GetID(),
		Titulum: q.GetTitulum(),
		Pars:    q.GetPars(),
		Hora:    q.GetHora(),
		Dies_id: q.GetDiesId(),
	}

	jsonStr, err := json.Marshal(newQuartum)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func (q *Quartum) CreateHourFromString(hora string) (time.Time, error) {

	// Define the expected format
	layout := "15:04:05"

	// Parse the string
	parsedTime, err := time.Parse(layout, hora)

	if err != nil {
		return time.Now(), err
	}

	return parsedTime, nil
}

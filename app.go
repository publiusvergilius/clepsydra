package main

import (
	"clepsydra/db"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Olá %s, It's show time!", name)
}

var prodDB db.DB
var r db.QuartumRepository
var dR db.DiesRepository

func init() {
	fmt.Println("---------Olá----------")
	sqlite, err := sql.Open("sqlite3", "./data.db") // ":memory:" or "./data.db
	if err != nil {
		log.Fatal(err)
	}
	// defer sqlite.Close()
	// sqlite.SetMaxOpenConns(1)

	r = db.QuartumRepository{}
	dR = db.DiesRepository{}

	prodDB = sqlite

	r.Create(prodDB)
	dR.Create(prodDB)

}

func (a *App) GetAllQuarta() string {
	quarta, err := r.FindAll(prodDB)
	if err != nil {
		return ""
	}

	var newQuarta []string

	for _, quartum := range quarta {
		str, err := quartum.ToString()
		if err == nil {

			newQuarta = append(newQuarta, str)
		}
	}

	return Stringfy(newQuarta)
}

func (a *App) CreateQuartum() error {
	d := db.Dies{}
	d.SetDate(time.Now())
	dR.Save(prodDB, d)

	q := db.Quartum{}
	q.SetTitulum("Clepsydra")
	q.SetPars(1)
	q.SetHora(time.Now())
	q.SetDiesId(1)

	err := r.Save(prodDB, q)
	if err != nil {
		return err
	}

	return nil
}

func Stringfy(jsonSlice []string) string {
	result := "[" + strings.Trim(strings.Join(jsonSlice, ","), " ") + "]"
	return result
}

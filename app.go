package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"clepsydra/db"
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

func (a *App) GetAllDies() string {
	diei, err := dR.FindAll(prodDB)
	if err != nil {
		log.Default().Fatalf("unexpected error occurred: %q", err)
		return "error"
	}

	fmt.Println(diei)
	var newDies []string
	for _, dies := range diei {
		str, err := dies.ToString()
		if err == nil {
			newDies = append(newDies, str)
		}
	}

	return Stringfy(newDies)
}

func (a *App) CreateDies(data string) string {
	type Request struct {
		Dies   string `json:"dies"`
		Format string `json:"format"`
	}
	var request Request

	fmt.Println(request)
	err := json.Unmarshal([]byte(data), &request)

	if err != nil {
		log.Fatal(err)
		return "error"
	}

	parsedTime, err := time.Parse("01/02/2006", request.Dies)
	if err != nil {
		log.Fatalf("Error parsing data: %q", err)
		return err.Error()
	}

	var dies db.Dies
	dies.SetDate(parsedTime)

	err = dR.Save(prodDB, dies)

	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return "created"

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

func (a *App) CreateQuartum(data string) string {
	var q db.Quartum
	err := json.Unmarshal([]byte(data), &q)

	if err != nil {
		log.Fatal(err)
		return "error"
	}

	err = r.Save(prodDB, q)

	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return "created"
}

func Stringfy(jsonSlice []string) string {
	result := "[" + strings.Trim(strings.Join(jsonSlice, ","), " ") + "]"
	return result
}

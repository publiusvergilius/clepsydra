package db

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// diretamente, o database é reiniciado a cada teste
/* Testing flow
* 1 - abre database
* 2 - alocação de tabelas
* 3 - alocação de dados nas tabelas
* 4 - CRUD
* 5 - fechar database
 */

var db = setupTestDB()

func TestRepository(t *testing.T) {

	diesRepo := DiesRepository{}
	_, err := diesRepo.newDiesRepository(db)
	testTableInitialization(t, "dies")

	if err != nil {
		t.Fatalf("Expected no error while creating dies repository, got %v", err)
	}

	quartumRepository := QuartumRepository{}
	_, err = quartumRepository.newQuartumRepository(db)
	testTableInitialization(t, "quartum")

	if err != nil {
		t.Fatalf("Expected no error while creating quartum repository, got %v", err)
	}

	defer db.Close()

	/* !TODO: parece estar testando muitas coisas ao mesmo tempo:
	*	 criação e verificação de tabelas; quantidade de rows alocados;
	*	formato de datas na inserção;
	* 1 - dia não pode ter valor anterior ao presente
	 */
	t.Run("insert row in dies and get it back", func(t *testing.T) {
		dies := Dies{}
		dies.SetDate(time.Now())

		db.Exec("insert into dies(date) values (?)",
			dies.GetDate())

		diesList := []Dies{}
		querySQL := "select id, date from dies"
		rows, err := db.Query(querySQL)
		if err != nil {
			t.Error("RAND query error on dies:", err)
		}
		defer rows.Close()

		for rows.Next() {
			var id uint
			var date string

			err := rows.Scan(&id, &date)
			if err != nil {
				t.Errorf("was not told to error: %q", err)
			}

			newDies := Dies{id: id}
			parsedDate, err := time.Parse(time.DateOnly, date)
			if err != nil {
				log.Fatalf("error while parsing date: %q", err)
			}
			newDies.SetDate(parsedDate)
			diesList = append(diesList, newDies)
		}

		want := Dies{1, time.Now().UTC().Format(time.DateOnly)}

		if len(diesList) != 1 {
			t.Errorf("list of dies should contain 1 element, got %q", len(diesList))
		}

		got := diesList[0]

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	/*!TODO:
	* 1 - trocar hora por finis(ou termino), initium
	* 2 - pars não pode ser menor que 0 ou maior que 4
	* 3 - fazer many-to-one com tabela do dies
	 */
	t.Run("insert row in quartum table and get it back", func(t *testing.T) {
		quartum := Quartum{}
		quartum.SetTitulum("programação")
		quartum.SetHora(time.Now())
		quartum.SetPars(1)
		fmt.Println(quartum.GetHora())

		db.Exec("insert into quartum(pars, titulum, hora) values (?, ?, ?);",
			quartum.GetPars(), quartum.GetTitulum(), quartum.GetHora())

		quartumList := []Quartum{}
		querySQL := "select id, titulum, hora, pars from quartum"
		rows, err := db.Query(querySQL)
		if err != nil {
			t.Error("RAND query error on quartum: ", err)
		}
		defer rows.Close()

		for rows.Next() {
			var id uint
			var titulum string
			var hora string
			var pars uint8

			err := rows.Scan(&id, &titulum, &hora, &pars)
			fmt.Println(hora)

			if err != nil {
				t.Errorf("was not told to error: %q", err)
			}

			newQuartum := Quartum{id: id, hora: hora}
			newQuartum.SetPars(pars)
			newQuartum.SetTitulum(titulum)
			// newQuartum.SetHora(hora)

			fmt.Println(newQuartum)

			quartumList = append(quartumList, newQuartum)
		}

		want := Quartum{
			id:      1,
			hora:    time.Now().UTC().Format(time.TimeOnly),
			titulum: "programação",
			pars:    1,
		}

		fmt.Println(want)

		if len(quartumList) != 1 {
			t.Errorf("list of dies should contain 1 element, got %q", len(quartumList))
		}

		got := quartumList[0]

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

}

func testTableInitialization(t *testing.T, tableName string) {
	t.Helper()

	var result string
	queryStr := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)

	err := db.QueryRow(queryStr).Scan(&result)

	if err == sql.ErrNoRows {
		t.Fatalf("Expected table %s to be created.", tableName)
	}

	if err != nil {
		t.Fatalf("Expected no error while verifying columns in the table.")
	}
}

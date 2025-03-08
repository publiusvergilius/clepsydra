package db

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/**
* 1 - verificar se Entity está sincronizado com a string que criar a tabela
* 2 - inserção incorreta levata erro sem com panic
 */

/**
* 1 - verificar se Entity está sincronizado com a string que criar a tabela
* 2 - inserção incorreta levata erro sem com panic
 */

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
	_, err := diesRepo.Create(db)
	testTableInitialization(t, "dies")

	if err != nil {
		t.Fatalf("Expected no error while creating dies repository, got %v", err)
	}

	quartumRepository := QuartumRepository{}
	_, err = quartumRepository.Create(db)
	testTableInitialization(t, "quartum")

	if err != nil {
		t.Fatalf("Expected no error while creating quartum repository, got %v", err)
	}

	defer db.Close()

	/** !TODO: parece estar testando muitas coisas ao mesmo tempo:
	*	 criação e verificação de tabelas; quantidade de rows alocados;
	*	formato de datas na inserção;
	* 1 - dia não pode ter valor anterior ao presente
	* 1 - trocar hora por finis(ou termino), initium
	* 2 - pars não pode ser menor que 0 ou maior que 4
	* 3 - fazer many-to-one com tabela do dies
	* 4 - insertion errors from sqlite3 arre not returned - checar erro de inserç
	* 4 - insertion errors from sqlite3 arre not returned - checar erro de inserç
	*/
	t.Run("insert row in quartum table and get it back", func(t *testing.T) {
		quartum := Quartum{id: 1}
		quartum.SetTitulum("programação")
		quartum.SetHora(time.Now())
		quartum.SetPars(1)
		quartum.SetDiesId(1)

		quartumRepository := QuartumRepository{}

		testOneEntityInsertion(t, db, quartumRepository, quartum)
	})

	type RepositoryCases[TR any, TE any] struct {
		Name           string
		Repository     TR
		ExpectedEntity TE
	}

	cases := []RepositoryCases[QuartumRepository, Quartum]{
		{
			"one entity operations",
			QuartumRepository{},
			Quartum{id: 1, titulum: "programação", pars: 1, dies_id: 1},
		},
	}

	/** testango encotrar por id*/
	// todo mover para assertFindById usando IoC
	// criar assertNotFindById
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			repository := test.Repository
			var id uint = 1
			got, err := repository.FindById(db, id)

			if err != nil {
				t.Error("unexpected error: ", err)
			}

			want := test.ExpectedEntity
			want.SetHora(time.Now())

			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}

		})
	}

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

func testOneEntityInsertion[T comparable](t *testing.T, db DB, repository Repository[T], want T) {

	err := repository.Save(db, want)

	if err != nil {
		t.Fatalf("was not told to error on saving to repository: %q", err)
	}

	list, err := repository.FindAll(db)

	if err != nil {
		t.Errorf("was not told to error on finding all: %q", err)
	}

	if len(list) != 1 {
		t.Fatalf("list of quartum should contain 1 element, got %q", len(list))
	}

	for _, item := range list {
		if item != want {
			t.Errorf("want %v, got %v", want, item)
		}
	}
}

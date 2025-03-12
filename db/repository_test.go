package db

import (
	"database/sql"
	"fmt"
	"slices"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
	defer db.Close()

	diesRepository := DiesRepository{}
	_, err := diesRepository.Create(db)

	testTableInitialization(t, "dies", err)

	quartumRepository := QuartumRepository{}
	_, err = quartumRepository.Create(db)

	testTableInitialization(t, "quartum", err)

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

	quartum_data := []Quartum{
		{id: 1, titulum: "programação", pars: 1, dies_id: 1},
		{id: 2, titulum: "música", pars: 2, dies_id: 1},
		{id: 3, titulum: "programação", pars: 1, dies_id: 2},
		{id: 4, titulum: "programação", pars: 1, dies_id: 2},
	}

	// insert first
	for i := 0; i < len(quartum_data); i++ {

		quartum := Quartum{id: quartum_data[i].GetID()}
		quartum.SetTitulum(quartum_data[i].GetTitulum())
		quartum.SetHora(time.Now())
		quartum.SetPars(quartum_data[i].GetPars())
		quartum.SetDiesId(quartum_data[i].GetDiesId())

		// insert and verify if insertion really happened
		testOneEntityInsertion(t, db, quartumRepository, quartum)

	}

	testTableNumberOfInsertions(t, db, uint(len(quartum_data)), quartumRepository)
	assertQuartumIsUnique(t, db, quartumRepository)

	type RepositoryCase[T Entity] struct {
		Name       string
		Repository Repository[T]
		Expected   []T
	}

	quartumRepoTest := RepositoryCase[Quartum]{
		Name:       "verify if inserted data exists in database",
		Repository: quartumRepository,
		Expected:   quartum_data,
	}

	t.Run(quartumRepoTest.Name, func(t *testing.T) {
		repository := quartumRepoTest.Repository

		for _, want := range quartumRepoTest.Expected {
			want.SetHora(time.Now())

			list, err := repository.FindAll(db)
			if err != nil {
				t.Error("was not told to error on find all quarta: ", err)
			}

			if !slices.Contains(list, want) {
				t.Errorf("%q was expected to exists in %q", want, list)
			}
		}

	})

	// assert quartum is not beeing repeated
	// assertQuartumIsUnique(t, db, quartumRepository)

	/** Get all quarta(plural) where dies_id is equal to dies.id*/

	/** create actio table */

}

func testTableInitialization(t *testing.T, tableName string, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Expected no error while creating dies repository, got %v", err)
	}

	var result string
	queryStr := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)

	err = db.QueryRow(queryStr).Scan(&result)

	if err == sql.ErrNoRows {
		t.Fatalf("Expected table %s to be created.", tableName)
	}

	if err != nil {
		t.Fatalf("Expected no error while verifying columns in the table.")
	}
}

func testOneEntityInsertion[T Entity](t *testing.T, db DB, repository Repository[T], want T) {
	t.Helper()

	err := repository.Save(db, want)

	if err != nil {
		t.Errorf("was not told to error on saving to repository: %q", err)
	}

	_, err = repository.FindById(db, want.GetID())
	fmt.Println(err)

	if err != nil {
		t.Fatalf("was not told to error on finding entity: %q", err)
	}
}

func testTableNumberOfInsertions[T Entity](t *testing.T, db DB, want uint, repository Repository[T]) {
	t.Helper()

	list, err := repository.FindAll(db)
	if err != nil {
		t.Fatalf("was not told to error")
	}

	if len(list) != int(want) {
		t.Errorf("list should contain %q element, got %q", want, len(list))
	}
}

func assertQuartumIsUnique(t *testing.T, db DB, repository Repository[Quartum]) {
	t.Helper()

	list, err := repository.FindAll(db)
	if err != nil {
		t.Fatalf("was not told to error on find all quarta")
	}

	if len(list) > 1 {
		for i, quartum := range list {
			for j, compared := range list {
				if i == j {
					continue
				}

				if quartum.GetTitulum() == compared.GetTitulum() && quartum.GetPars() == compared.GetPars() && quartum.GetDiesId() == compared.GetDiesId() {
					t.Errorf("was not told to have repeated quarta: %q", quartum)
				}
			}
		}
	}
}

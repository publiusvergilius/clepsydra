package db

import (
	"database/sql"
	"fmt"
	"reflect"
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
		{id: 1, titulum: "programação", pars: 1, dies_id: 1, prazo: "10:10:10"},
		{id: 2, titulum: "música", pars: 2, dies_id: 1},
		{id: 3, titulum: "programação", pars: 1, dies_id: 2},
		{id: 4, titulum: "programação", pars: 2, dies_id: 2},
	}

	// insert first
	for _, q := range quartum_data {

		quartum := Quartum{id: q.GetID()}
		quartum.SetTitulum(q.GetTitulum())
		quartum.SetHora(time.Now())
		quartum.SetPars(q.GetPars())
		quartum.SetDiesId(q.GetDiesId())

		// insert and verify if insertion really happened

		want := q.prazo
		parsedTime, _ := time.Parse("15:10:00", want)
		q.SetPrazo(parsedTime)
		got := quartum.GetPrazo()

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}

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

	quartumRepoTest = RepositoryCase[Quartum]{
		Name:       "verify if first element was deleted",
		Repository: quartumRepository,
		Expected:   quartum_data[1:],
	}

	quartumRepository.Delete(db, 1)
	assertNotExist(t, db, quartumRepository, 1)
	testTableNumberOfInsertions(t, db, uint(len(quartum_data)-1), quartumRepository)

	// assert quartum is not beeing repeated
	// assertQuartumIsUnique(t, db, quartumRepository)

	/** Get all quarta(plural) where dies_id is equal to dies.id*/

	/** create actio table */

	/** Where dies get it's quarta */

	diesRepoTest := RepositoryCase[Dies]{
		Name: "find all quarta that references dies",
	}

	var dies_id uint = 1

	got, _ := diesRepository.FindQuarta(db, dies_id)

	/** to be compared with dies_quarta */
	t.Run(diesRepoTest.Name, func(t *testing.T) {
		quartumList, err := quartumRepository.FindAll(db)

		var want []Quartum
		if err != nil {
			t.Fatal("was not told to error: ", err)
		}

		for _, quartum := range quartumList {
			if quartum.GetDiesId() == dies_id {
				want = append(want, quartum)
			}

		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %q, got %q", want, got)
		}
	})

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

func assertExist[T Entity](t *testing.T, db DB, repository Repository[T], id uint) {
	_, err := repository.FindById(db, id)

	if err != nil {
		t.Fatalf("was not told to error on finding element with %q: %q", id, err)
	}

}

func assertNotExist[T Entity](t *testing.T, db DB, repository Repository[T], id uint) {
	element, err := repository.FindById(db, id)

	if err == nil {
		t.Fatalf("was told to error on finding element with %q", element.GetID())
	}

}

func testOneEntityInsertion[T Entity](t *testing.T, db DB, repository Repository[T], want T) {
	t.Helper()

	err := repository.Save(db, want)

	if err != nil {
		t.Errorf("was not told to error on saving to repository: %q", err)
	}

	_, err = repository.FindById(db, want.GetID())

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

func assertTimestamp(t *testing.T, want string) {

	t.Helper()
	parsedTime, err := time.Parse("15:04:05", want)

	if err != nil {
		t.Errorf("was not told to error on parsing time: %q\n", err)
	}

	got := parsedTime.Format("12:10:00")

	if got != want {
		t.Errorf("wanted %q, got %q\n", want, got)
	}
}

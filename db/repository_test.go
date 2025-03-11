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

	diesRepo := DiesRepository{}
	_, err := diesRepo.Create(db)

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
		{id: 4, titulum: "programação", pars: 2, dies_id: 1},
	}

	// insert first
	for i := 0; i < len(quartum_data); i++ {

		quartum := Quartum{id: quartum_data[i].GetID()}
		quartum.SetTitulum(quartum_data[i].GetTitulum())
		quartum.SetHora(time.Now())
		quartum.SetPars(quartum_data[i].GetPars())
		quartum.SetDiesId(quartum_data[i].GetDiesId())

		// verify if insertion really happened
		testOneEntityInsertion(t, db, quartumRepository, quartum)

	}

	testTableNumberOfInsertions(t, db, uint(len(quartum_data)), quartumRepository)

	type RepositoryCases struct {
		Name       string
		Repository Repository[Quartum]
		Expected   []Quartum
	}

	cases := []RepositoryCases{
		{
			Name:       "one quartum find operation",
			Repository: quartumRepository,
			Expected: []Quartum{
				{id: 1, titulum: "programação", pars: 1, dies_id: 1},
			},
		},
		{
			Name:       "multiple quartum find operations",
			Repository: quartumRepository,
			Expected:   quartum_data,
		},
	}

	/** testando encotrar por id*/
	// todo mover para assertFindById usando IoC
	// criar assertNotFindById
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			// count number of insertions
			// testTableNumberOfInsertions(t, db, uint(i+1), quartumRepository)

			repository := test.Repository
			var id uint = 1
			got, err := repository.FindById(db, id)

			if err != nil {
				t.Error("unexpected error: ", err)
			}

			want := test.Expected

			row := want[0]
			row.SetHora(time.Now())
			if row != got {
				t.Errorf("want %q, got %q", want, got)
			}

		})

	}

	// assert data is not repeated
	assertUniqueTitulumPars(t, db, quartumRepository)

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
		t.Fatalf("was not told to error on saving to repository: %q", err)
	}

	_, err = repository.FindById(db, want.GetID())

	if err != nil {
		t.Errorf("was not told to error on finding entity: %q", err)
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

func assertUniqueTitulumPars(t *testing.T, db DB, repository Repository[Quartum]) {
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

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

var ent Entity = Quartum{}

type Repo[T Entity] interface {
	FindAll(db DB) ([]T, error)
}

type QuartumRepo struct{}

func (QuartumRepo) FindAll(db DB) ([]Quartum, error) {
	return []Quartum{{}}, nil
}

func testFindAll[T Entity](r Repo[T], db DB) {
	r.FindAll(db)
}

var db = setupTestDB()

func TestRepository(t *testing.T) {
	defer db.Close()

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

	})

	type RepositoryCases struct {
		Name       string
		Repository Repository[Quartum]
		Expected   []Quartum
	}

	cases := []RepositoryCases{
		{
			Name:       "one quartum find operation",
			Repository: QuartumRepository{},
			Expected: []Quartum{
				{id: 1, titulum: "programação", pars: 1, dies_id: 1},
			},
		},
		{
			Name:       "multiple quartum find operations",
			Repository: QuartumRepository{},
			Expected: []Quartum{
				{id: 1, titulum: "programação", pars: 1, dies_id: 1},
				{id: 2, titulum: "música", pars: 2, dies_id: 1},
			},
		},
	}

	/** testando encotrar por id*/
	// todo mover para assertFindById usando IoC
	// criar assertNotFindById
	for i, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			// insert first
			quartum := Quartum{id: 1}
			quartum.SetTitulum("programação")
			quartum.SetHora(time.Now())
			quartum.SetPars(1)
			quartum.SetDiesId(1)

			testOneEntityInsertion(t, db, quartumRepository, quartum)
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

		if i != 0 {
			// insert second
			quartum := Quartum{id: 2}
			quartum.SetTitulum("música")
			quartum.SetHora(time.Now())
			quartum.SetPars(2)
			quartum.SetDiesId(1)

			testOneEntityInsertion(t, db, quartumRepository, quartum)
			testTableNumberOfInsertions(t, db, 2, quartumRepository)

			t.Run(test.Name, func(t *testing.T) {
				repository := test.Repository
				var id uint = 1
				got, err := repository.FindById(db, id)

				if err != nil {
					t.Error("unexpected error: ", err)
				}

				want := test.Expected

				for _, row := range want {
					row.SetHora(time.Now())
					if row != got {
						t.Errorf("want %q, got %q", want, got)
					}
				}

			})
		}
	}

	/** Get all quarta(plural) where dies_id is equal to dies.id*/

	/** create actio table */

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

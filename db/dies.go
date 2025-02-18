package db

type Dies struct {
	ID      uint
	Titulum string
	Date    string // 31-01-2000
}

type DiesRepository struct{}

func NewDiesRepository(db DB) (Result, error) {
	sqlStmt := `create table dies (id integer not null primary key, titulum text);`

	return db.Exec(sqlStmt)
}

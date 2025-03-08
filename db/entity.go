package db

type Entity[T any] any

type Repository[T comparable] interface {
	FindAll(db DB) ([]T, error)
	FindById(db DB, id uint) (T, error)
	Save(db DB, entity T) error
	Create(db DB) (Result, error)
}

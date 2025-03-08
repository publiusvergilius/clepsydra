package db

type Entity[T any] any

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindById(id uint) (T, error)
	Save(entity T) error
	Create(entity T) error
}

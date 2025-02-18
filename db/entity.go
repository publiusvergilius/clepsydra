package db

type Entity[T any] interface {
	GetAll() ([]T, error)
	GetById(id uint) (T, error)
	Save(entity T) error
}

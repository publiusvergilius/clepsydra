package db

type Repository[T Entity] interface {
	FindAll(db DB) ([]T, error)
	FindById(db DB, id uint) (T, error)
	Save(db DB, entity T) error
	Create(db DB) (Result, error)
}

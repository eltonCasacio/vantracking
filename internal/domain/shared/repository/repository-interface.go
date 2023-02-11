package repository

type RepositoryInterface[T interface{}] interface {
	Create(aggragate *T) error
	Update(aggragate *T) error
	Delete(id string) error
	FindAll() ([]T, error)
	FindByID(id string) (*T, error)
}

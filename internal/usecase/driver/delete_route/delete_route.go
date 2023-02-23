package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type DeleteRouteUseCase struct {
	repository repo.DriverRepositoryInterface
}

func NewUseCase(repository repo.DriverRepositoryInterface) *DeleteRouteUseCase {
	return &DeleteRouteUseCase{
		repository: repository,
	}
}

func (cd *DeleteRouteUseCase) DeleleRoute(id string) error {
	err := cd.repository.DeleteRoute(id)
	if err != nil {
		return err
	}
	return nil
}

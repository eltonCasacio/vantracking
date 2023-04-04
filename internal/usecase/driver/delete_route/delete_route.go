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

func (cd *DeleteRouteUseCase) DeleleRoute(code string) error {
	err := cd.repository.DeleteRoute(code)
	if err != nil {
		return err
	}
	return nil
}

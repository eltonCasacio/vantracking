package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type DeleteDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *DeleteDriverUseCase {
	return &DeleteDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (u *DeleteDriverUseCase) Delete(id string) error {
	err := u.driverRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

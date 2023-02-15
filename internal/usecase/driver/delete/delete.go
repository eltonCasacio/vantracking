package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type deleteDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *deleteDriverUseCase {
	return &deleteDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (u *deleteDriverUseCase) Delete(id string) error {
	err := u.driverRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

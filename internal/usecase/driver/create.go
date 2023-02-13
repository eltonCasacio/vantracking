package usecase

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type createDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func CreateDriverUseCase(driverRepository repo.DriverRepositoryInterface) *createDriverUseCase {
	return &createDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *createDriverUseCase) Execute(input f.DriverInputDTO) error {
	driver, err := f.DriverFactory().Create(input)
	if err != nil {
		return err
	}

	err = cd.driverRepository.Create(driver)
	if err != nil {
		return err
	}
	return nil
}

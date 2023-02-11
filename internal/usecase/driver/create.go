package usecase

import (
	d "github.com/eltoncasacio/vantracking/internal/domain/driver"
)

type createDriverUseCase struct {
	criverRepository d.DriverRepositoryInterface
}

func CreateDriverUseCase(driverRepository d.DriverRepositoryInterface) *createDriverUseCase {
	return &createDriverUseCase{
		criverRepository: driverRepository,
	}
}

func (cd *createDriverUseCase) Execute(input d.DriverInputDTO) error {
	driver, err := d.DriverFactory().Create(input)
	if err != nil {
		return err
	}

	err = cd.criverRepository.Create(driver)
	if err != nil {
		return err
	}
	return nil
}

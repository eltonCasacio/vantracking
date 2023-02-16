package driver

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type RegisterDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *RegisterDriverUseCase {
	return &RegisterDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *RegisterDriverUseCase) RegisterDriver(input DriverInputDTO) error {
	driverInput := f.CreateDriverInputDTO{
		ID:       "",
		CPF:      input.CPF,
		Name:     input.Name,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		UF:       input.UF,
		City:     input.City,
		Street:   input.Street,
		Number:   input.Number,
		CEP:      input.CEP,
	}

	driver, err := f.DriverFactory().Create(driverInput)
	if err != nil {
		return err
	}

	err = cd.driverRepository.Create(driver)
	if err != nil {
		return err
	}
	return nil
}

package driver

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

func (cd *createDriverUseCase) Execute(input DriverInputDTO) error {
	data := f.DriverInputDTO{
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
	driver, err := f.DriverFactory().Create(data)
	if err != nil {
		return err
	}

	err = cd.driverRepository.Create(driver)
	if err != nil {
		return err
	}
	return nil
}

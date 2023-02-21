package driver

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type updateDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *updateDriverUseCase {
	return &updateDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (u *updateDriverUseCase) Update(input DriverInputDTO) error {
	driverInput := f.CreateInstanceDriverInputDTO{
		ID:       input.ID,
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

	driver, err := f.DriverFactory().CreateInstance(driverInput)
	if err != nil {
		return err
	}

	err = u.driverRepository.Update(driver)
	if err != nil {
		return err
	}
	return nil
}

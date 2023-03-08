package driver

import (
	"errors"

	f "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type RegisterDriverUseCase struct {
	repository repo.DriverRepositoryInterface
}

func NewUseCase(repository repo.DriverRepositoryInterface) *RegisterDriverUseCase {
	return &RegisterDriverUseCase{
		repository: repository,
	}
}

func (u *RegisterDriverUseCase) RegisterDriver(input DriverInputDTO) error {
	driverInput := f.NewDriverInputDTO{
		CPF:    input.CPF,
		Name:   input.Name,
		Phone:  input.Phone,
		UF:     input.UF,
		City:   input.City,
		Street: input.Street,
		Number: input.Number,
		CEP:    input.CEP,
	}

	driver, err := f.DriverFactory().New(driverInput)
	if err != nil {
		return err
	}

	found, _ := u.repository.FindByCPF(input.CPF)
	if found != nil {
		return errors.New("driver already exists")
	}

	err = u.repository.Create(driver)
	if err != nil {
		return err
	}
	return nil
}

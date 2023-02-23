package driver

import (
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

// func (cd *RegisterDriverUseCase) RegisterDriver(input DriverInputDTO) error {
// 	driverInput := f.NewDriverInputDTO{
// 		CPF:    input.CPF,
// 		Name:   input.Name,
// 		Phone:  input.Phone,
// 		UF:     input.UF,
// 		City:   input.City,
// 		Street: input.Street,
// 		Number: input.Number,
// 		CEP:    input.CEP,
// 	}

// 	driver, err := f.DriverFactory().New(driverInput)
// 	if err != nil {
// 		return err
// 	}

// 	found, _ := cd.driverRepository.FindByCPF(input.CPF)
// 	if found != nil {
// 		return errors.New("driver already exists")
// 	}

// 	err = cd.driverRepository.Create(driver)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

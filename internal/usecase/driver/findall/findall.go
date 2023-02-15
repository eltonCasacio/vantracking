package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type findAllDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *findAllDriverUseCase {
	return &findAllDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *findAllDriverUseCase) FindAll() ([]DriverOutputDTO, error) {
	driversFound, err := cd.driverRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var drivers []DriverOutputDTO
	for _, driver := range driversFound {
		addr := driver.GetAddress()
		output := DriverOutputDTO{
			ID:       driver.GetID().String(),
			CPF:      driver.GetCPF(),
			Name:     driver.GetName(),
			Nickname: driver.GetNickName(),
			Phone:    driver.GetPhone(),
			UF:       addr.GetUF(),
			City:     addr.GetCity(),
			Street:   addr.GetStreet(),
			Number:   addr.GetNumber(),
			CEP:      addr.GetCEP(),
		}
		drivers = append(drivers, output)
	}

	return drivers, nil
}

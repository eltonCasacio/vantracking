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

func (cd *findAllDriverUseCase) ListAll() ([]DriverOutputDTO, error) {
	driversFound, err := cd.driverRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var drivers []DriverOutputDTO
	for _, driver := range driversFound {
		addr := driver.Address
		output := DriverOutputDTO{
			ID:       driver.ID.String(),
			CPF:      driver.CPF,
			Name:     driver.Name,
			Nickname: driver.Nickname,
			Phone:    driver.Phone,
			UF:       addr.UF,
			City:     addr.City,
			Street:   addr.Street,
			Number:   addr.Number,
			CEP:      addr.CEP,
		}
		drivers = append(drivers, output)
	}

	return drivers, nil
}

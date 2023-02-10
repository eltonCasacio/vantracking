package usecase

import (
	a "github.com/eltoncasacio/vanmonit/internal/domain/address"
	entity "github.com/eltoncasacio/vanmonit/internal/domain/driver"
	s "github.com/eltoncasacio/vanmonit/internal/domain/school"
)

type CreateDriverUseCase struct {
	DriverRepository entity.DriverRepositoryInterface
}

func NewCreateDriverUseCase(driverRepository entity.DriverRepositoryInterface) *CreateDriverUseCase {
	return &CreateDriverUseCase{
		DriverRepository: driverRepository,
	}
}

func (cd *CreateDriverUseCase) Execute(input DriverInputDTO) error {
	addrSchool, err := a.NewAddresses(input.UFSchool, input.CitySchool, input.StreetSchool, input.NumberSchool, input.CEPSchool)
	if err != nil {
		return err
	}
	addrDriver, err := a.NewAddresses(input.UFAddress, input.CityAddress, input.StreetAddress, input.NumberAddress, input.CEPAddress)
	if err != nil {
		return err
	}

	school, err := s.NewSchool(input.SchoolName, addrSchool)
	if err != nil {
		return err
	}

	driver, err := entity.NewDriver(input.CPF, input.Name, input.Nickname, input.Phone, input.PlateNumber, addrDriver, school)
	if err != nil {
		return err
	}

	err = cd.DriverRepository.Create(driver)
	if err != nil {
		return err
	}

	return nil
}

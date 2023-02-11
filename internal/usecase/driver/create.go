package usecase

import (
	entity "github.com/eltoncasacio/vantracking/internal/domain/entities/driver"
	vo "github.com/eltoncasacio/vantracking/internal/domain/value_objects"
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
	addrSchool, err := vo.NewAddresses(input.UFSchool, input.CitySchool, input.StreetSchool, input.NumberSchool, input.CEPSchool)
	if err != nil {
		return err
	}
	addrDriver, err := vo.NewAddresses(input.UFAddress, input.CityAddress, input.StreetAddress, input.NumberAddress, input.CEPAddress)
	if err != nil {
		return err
	}

	school, err := vo.NewSchool(input.SchoolName, addrSchool)
	if err != nil {
		return err
	}

	schools := []vo.School{*school}

	driver, err := entity.NewDriver(input.CPF, input.Name, input.Nickname, input.Phone, input.PlateNumber, schools, *addrDriver)
	if err != nil {
		return err
	}

	err = cd.DriverRepository.Create(driver)
	if err != nil {
		return err
	}

	return nil
}

package factory

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type driverFactory struct{}

func DriverFactory() *driverFactory {
	return &driverFactory{}
}

func (df *driverFactory) New(input NewDriverInputDTO) (*e.Driver, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP, input.Complement, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}

	driver, err := e.NewDriver(input.CPF, input.Name, input.Phone, input.Nickname, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (df *driverFactory) CreateInstance(input CreateInstanceDriverInputDTO) (*e.Driver, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP, input.Complement, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}

	id, err := identity.ParseID(input.ID)
	if err != nil {
		return nil, err
	}

	driver := e.Driver{
		ID:       id,
		CPF:      input.CPF,
		Name:     input.Name,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		Address:  *addrDriver,
	}

	if err := driver.IsValid(); err != nil {
		return nil, err
	}

	return &driver, nil
}

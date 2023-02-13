package factory

import (
	"sync"

	e "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type DriverInputDTO struct {
	CPF      string
	Name     string
	Nickname string
	Phone    string
	UF       string
	City     string
	Street   string
	Number   string
	CEP      int
}

type driverFactory struct{}

var instance *driverFactory
var lock = &sync.Mutex{}

func DriverFactory() *driverFactory {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &driverFactory{}
		}
	}
	return instance
}

func (df *driverFactory) Create(input DriverInputDTO) (*e.Driver, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP)
	if err != nil {
		return nil, err
	}

	driver, err := e.NewDriver(input.CPF, input.Name, input.Nickname, input.Phone, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
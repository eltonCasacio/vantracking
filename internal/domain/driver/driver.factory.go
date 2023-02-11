package driver

import (
	"sync"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/value_objects"
)

type driverFactory struct {
}

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

func (df *driverFactory) Create(input DriverInputDTO) (*Driver, error) {
	addrSchool, err := vo.NewAddresses(input.UFSchool, input.CitySchool, input.StreetSchool, input.NumberSchool, input.CEPSchool)
	if err != nil {
		return nil, err
	}
	addrDriver, err := vo.NewAddresses(input.UFAddress, input.CityAddress, input.StreetAddress, input.NumberAddress, input.CEPAddress)
	if err != nil {
		return nil, err
	}

	school, err := vo.NewSchool(input.SchoolName, addrSchool)
	if err != nil {
		return nil, err
	}

	schools := []vo.School{*school}

	driver, err := newDriver(input.CPF, input.Name, input.Nickname, input.Phone, input.PlateNumber, schools, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

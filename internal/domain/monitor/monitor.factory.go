package monitor

import (
	"sync"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type MonitorInputDTO struct {
	name        string
	cpf         string
	phoneNumber string
	UF          string
	City        string
	Street      string
	Number      string
	CEP         int
}

type monitorFactory struct{}

var instance *monitorFactory
var lock = &sync.Mutex{}

func MonitorFactory() *monitorFactory {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &monitorFactory{}
		}
	}
	return instance
}

func (df *monitorFactory) Create(input MonitorInputDTO) (*Monitor, error) {
	addrDriver, err := vo.NewAddresses(input.UF, input.City, input.Street, input.Number, input.CEP)
	if err != nil {
		return nil, err
	}

	driver, err := newMonitor(input.name, input.cpf, input.phoneNumber, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

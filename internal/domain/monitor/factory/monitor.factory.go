package factory

import (
	"sync"

	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

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

func (df *monitorFactory) Create(input CreateMonitorInputDTO) (*entity.Monitor, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP)
	if err != nil {
		return nil, err
	}

	driver, err := entity.NewMonitor(input.Name, input.CPF, input.PhoneNumber, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

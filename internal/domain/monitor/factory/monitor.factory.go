package factory

import (
	"sync"

	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
	vo_Passenger "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"

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
	CEP         string
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

func (df *monitorFactory) Create(input MonitorInputDTO) (*entity.Monitor, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP)
	if err != nil {
		return nil, err
	}

	driver, err := entity.NewMonitor(input.name, input.cpf, input.phoneNumber, *addrDriver, []vo_Passenger.Passenger{})
	if err != nil {
		return nil, err
	}

	return driver, nil
}

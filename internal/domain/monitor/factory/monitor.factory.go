package factory

import (
	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
	"github.com/eltoncasacio/vantracking/pkg/identity"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type monitorFactory struct{}

func MonitorFactory() *monitorFactory {
	return &monitorFactory{}
}

func (df *monitorFactory) Create(input NewMonitorInputDTO) (*entity.Monitor, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP, input.Complement, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}

	driver, err := entity.NewMonitor(input.Name, input.CPF, input.PhoneNumber, *addrDriver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (df *monitorFactory) Instance(input InstanceMonitorInputDTO) (*entity.Monitor, error) {
	addrDriver, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP, input.Complement, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}

	id, _ := identity.ParseID(input.ID)
	driver := entity.Monitor{
		ID:          id,
		Name:        input.Name,
		CPF:         input.CPF,
		PhoneNumber: input.PhoneNumber,
		Address:     *addrDriver,
	}

	return &driver, nil
}

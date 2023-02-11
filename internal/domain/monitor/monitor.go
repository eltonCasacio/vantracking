package monitor

import (
	"github.com/eltoncasacio/vantracking/pkg/entity"
)

type Monitor struct {
	id                entity.ID
	name              string
	cpf               string
	phoneNumber       string
	passengerName     string
	passengerNickname string
	schoolCode        string
	driverCPF         string
}

func NewMonitor(name, cpf, phoneNumber, passengerName, passengerNickname, schoolCode, driverCPF string) (*Monitor, error) {
	m := &Monitor{
		id:                entity.NewID(),
		name:              name,
		cpf:               cpf,
		phoneNumber:       phoneNumber,
		passengerName:     passengerName,
		passengerNickname: passengerNickname,
		schoolCode:        schoolCode,
		driverCPF:         driverCPF,
	}

	err := m.IsValid()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Monitor) IsValid() error {
	return nil
}

func (m *Monitor) GetID() entity.ID {
	return m.id
}

func (m *Monitor) GetName() string {
	return m.name
}

func (m *Monitor) GetCPF() string {
	return m.cpf
}

func (m *Monitor) GetPhoneNumber() string {
	return m.phoneNumber
}

func (m *Monitor) GetPassengerName() string {
	return m.passengerName
}

func (m *Monitor) GetPassengerNickname() string {
	return m.passengerNickname
}

func (m *Monitor) GetSchoolCode() string {
	return m.schoolCode
}

func (m *Monitor) GetDriverCPF() string {
	return m.driverCPF
}

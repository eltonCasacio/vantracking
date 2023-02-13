package monitor

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/entity"
)

type Monitor struct {
	id          entity.ID
	name        string
	cpf         string
	phoneNumber string
	address     vo.Address
	passengers  []Passenger
}

func newMonitor(name, cpf, phoneNumber string, address vo.Address) (*Monitor, error) {
	m := &Monitor{
		id:          entity.NewID(),
		name:        name,
		cpf:         cpf,
		phoneNumber: phoneNumber,
		address:     address,
	}

	err := m.IsValid()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Monitor) IsValid() error {
	var errs error
	if err := m.name == ""; err {
		return errors.New("name is required")
	}
	if err := m.cpf == ""; err {
		return errors.New("cpf is required")
	}
	if err := m.address.IsValid() != nil; err {
		return errors.New("address must be a valid address")
	}
	return errs
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

func (m *Monitor) AddPassenger(passenger Passenger) error {
	if err := passenger.IsValid(); err != nil {
		return err
	}
	m.passengers = append(m.passengers, passenger)
	return nil
}

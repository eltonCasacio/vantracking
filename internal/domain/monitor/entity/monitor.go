package entity

import (
	"errors"

	pEntity "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Monitor struct {
	id          identity.ID
	name        string
	cpf         string
	phoneNumber string
	address     vo.Address
	passengers  []pEntity.Passenger
}

func NewMonitor(name, cpf, phoneNumber string, address vo.Address, passengers []pEntity.Passenger) (*Monitor, error) {
	m := &Monitor{
		id:          identity.NewID(),
		name:        name,
		cpf:         cpf,
		phoneNumber: phoneNumber,
		address:     address,
		passengers:  []pEntity.Passenger{},
	}

	err := m.IsValid()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Monitor) IsValid() error {
	var errs error
	if err := m.GetName() == ""; err {
		return errors.New("invalid name")
	}
	if err := m.GetCPF() == ""; err {
		return errors.New("invalid cpf")
	}
	if err := m.GetPhoneNumber() == ""; err {
		return errors.New("invalid phonenumber")
	}
	if err := m.address.IsValid() != nil; err {
		return errors.New("address must be a valid address")
	}
	return errs
}

func (m *Monitor) GetID() identity.ID {
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

func (m *Monitor) GetAddress() vo.Address {
	return m.address
}

func (m *Monitor) GetPassengers() []pEntity.Passenger {
	return m.passengers
}

func (m *Monitor) AddPassenger(passenger pEntity.Passenger) error {
	if err := passenger.IsValid(); err != nil {
		return err
	}
	m.passengers = append(m.passengers, passenger)
	return nil
}

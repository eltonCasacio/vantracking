package entity

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Monitor struct {
	ID          identity.ID
	Name        string
	CPF         string
	PhoneNumber string
	Address     vo.Address
}

func NewMonitor(name, cpf, phoneNumber string, address vo.Address) (*Monitor, error) {
	m := &Monitor{
		ID:          identity.NewID(),
		Name:        name,
		CPF:         cpf,
		PhoneNumber: phoneNumber,
		Address:     address,
	}

	err := m.IsValid()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Monitor) IsValid() error {
	var errs error
	if err := m.Name == ""; err {
		return errors.New("invalid name")
	}
	if err := m.CPF == ""; err {
		return errors.New("invalid cpf")
	}
	if err := m.PhoneNumber == ""; err {
		return errors.New("invalid phonenumber")
	}
	if err := m.Address.IsValid() != nil; err {
		return errors.New("address must be a valid address")
	}
	return errs
}

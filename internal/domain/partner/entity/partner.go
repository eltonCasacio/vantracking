package entity

import (
	"errors"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Partner struct {
	ID          identity.ID
	Name        string
	Description string
	Price       float64
	PhoneNumber string
	Address     valueobjects.Address
	CategoryID  identity.ID
}

func NewPartner(name, description, phoneNumber string, address valueobjects.Address, categoryID identity.ID) (*Partner, error) {
	return &Partner{
		ID:          identity.NewID(),
		Name:        name,
		Description: description,
		PhoneNumber: phoneNumber,
		Address:     address,
		CategoryID:  categoryID,
	}, nil
}

func (p *Partner) IsValid() error {
	var errs error
	if err := p.Name == ""; err {
		return errors.New("name is required")
	}
	if err := p.Description == ""; err {
		return errors.New("description is required")
	}
	if err := p.PhoneNumber == ""; err {
		return errors.New("phone is required")
	}
	if err := p.Address.IsValid(); err != nil {
		return errors.New("address is required")
	}
	if p.CategoryID.String() == "" {
		return errors.New("id category is required")
	}
	return errs
}

package entity

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type PartnerCategory struct {
	ID   identity.ID
	Name string
}

func NewPartnerCategory(name string) (*PartnerCategory, error) {
	return &PartnerCategory{
		ID:   identity.NewID(),
		Name: name,
	}, nil
}

func (a *PartnerCategory) IsValid() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

package factory

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/partner/entity"
	"github.com/eltoncasacio/vantracking/pkg/identity"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type partnerFactory struct{}

func NewPartnerFactory() *partnerFactory {
	return &partnerFactory{}
}

func (df *partnerFactory) NewInstance(input PartnerInput) (*e.Partner, error) {
	addrPartner, err := vo.NewAddress(input.UF, input.City, input.Street, input.Number, input.CEP, input.Complement, input.Latitude, input.Longitude)
	if err != nil {
		return nil, err
	}

	id, err := identity.ParseID(input.ID)
	if err != nil {
		id = identity.NewID()
	}

	return &e.Partner{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		PhoneNumber: input.PhoneNumber,
		Address:     *addrPartner,
	}, nil
}

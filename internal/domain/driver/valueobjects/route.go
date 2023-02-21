package valueobjects

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Route struct {
	Code     string
	Name     string
	Origin   vo.Address
	Destiny  vo.Address
	DriverID identity.ID
}

func NewRoute(driverID identity.ID, code, name string, origin vo.Address, destiny vo.Address) (*Route, error) {
	r := &Route{
		DriverID: driverID,
		Code:     code,
		Name:     name,
		Origin:   origin,
		Destiny:  destiny,
	}
	err := r.IsValid()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Route) IsValid() error {
	if err := r.Code == ""; err {
		return errors.New("invalid code")
	}
	if err := r.Name == ""; err {
		return errors.New("invalid name")
	}
	if err := r.Origin.IsValid(); err != nil {
		return errors.New("origin is required")
	}
	if err := r.Destiny.IsValid(); err != nil {
		return errors.New("destiny is required")
	}
	return nil
}

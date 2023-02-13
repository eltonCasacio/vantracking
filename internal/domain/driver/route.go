package driver

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/entity"
)

type Route struct {
	code     string
	name     string
	origin   vo.Address
	destiny  vo.Address
	driverID entity.ID
}

func NewRoute(code, name string, origin vo.Address, destiny vo.Address) (*Route, error) {
	r := &Route{
		code:    code,
		name:    name,
		origin:  origin,
		destiny: destiny,
	}
	err := r.IsValid()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Route) IsValid() error {
	if err := r.driverID.String() == ""; err {
		return errors.New("driver id is required")
	}
	if err := r.name == ""; err {
		return errors.New("name id is required")
	}
	if err := r.origin.IsValid(); err != nil {
		return errors.New("origin is required")
	}
	if err := r.destiny.IsValid(); err != nil {
		return errors.New("destiny is required")
	}
	return nil
}

func (r *Route) GetCode() string {
	return r.code
}

func (r *Route) GetName() string {
	return r.name
}

func (r *Route) GetOrigin() vo.Address {
	return r.origin
}

func (r *Route) GetDestiny() vo.Address {
	return r.destiny
}

func (r *Route) GetDriverID() entity.ID {
	return r.driverID
}

package route

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Route struct {
	Code     identity.ID
	Name     string
	DriverID identity.ID
	Started  bool
}

func NewRoute(driverID identity.ID, name string) (*Route, error) {
	r := &Route{
		DriverID: driverID,
		Code:     identity.NewID(),
		Name:     name,
		Started:  false,
	}
	err := r.IsValid()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Route) IsValid() error {
	if err := r.Name == ""; err {
		return errors.New("invalid name")
	}
	return nil
}

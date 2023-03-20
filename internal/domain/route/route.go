package route

import (
	"errors"
	"strings"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Route struct {
	Code     string
	Name     string
	DriverID identity.ID
	Started  bool
}

func NewRoute(driverID identity.ID, name string) (*Route, error) {
	id := identity.NewID().String()
	splitID := strings.Split(id, "-")
	code := splitID[len(splitID)-1]
	r := &Route{
		DriverID: driverID,
		Code:     code,
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

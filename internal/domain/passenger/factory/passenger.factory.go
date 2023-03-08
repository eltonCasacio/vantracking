package passenger

import (
	"errors"
	"sync"

	e "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type passengerFactory struct{}

var instance *passengerFactory
var lock = &sync.Mutex{}

func PassengerFactory() *passengerFactory {
	return &passengerFactory{}
}

func (df *passengerFactory) NewPassenger(input NewPassengerInputDTO) (*e.Passenger, error) {
	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return nil, err
	}
	p, err := e.NewPassenger(
		input.Name,
		input.RouteCode,
		input.Nickname,
		input.SchoolName,
		monitorID,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (df *passengerFactory) Instance(input PassengerInputDTO) (*e.Passenger, error) {
	id, err := identity.ParseID(input.ID)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return nil, errors.New("invalid monitor id")
	}

	p := e.Passenger{
		ID:                id,
		Name:              input.Name,
		Nickname:          input.Nickname,
		RouteCode:         input.RouteCode,
		Goes:              input.Goes,
		Comesback:         input.Comesback,
		RegisterConfirmed: input.RegisterConfirmed,
		SchoolName:        input.SchoolName,
		MonitorID:         monitorID,
	}
	if err := p.IsValid(); err != nil {
		return nil, err
	}
	return &p, nil
}

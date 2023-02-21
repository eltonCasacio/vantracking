package passenger

import (
	"sync"

	e "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type passengerFactory struct{}

var instance *passengerFactory
var lock = &sync.Mutex{}

func PassengerFactory() *passengerFactory {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &passengerFactory{}
		}
	}
	return instance
}

func (df *passengerFactory) New(input PassengerInputDTO) (*e.Passenger, error) {
	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return nil, err
	}
	p, err := e.NewPassenger(
		input.Name,
		input.RouteCode,
		monitorID,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (df *passengerFactory) CreateInstance(input PassengerInputDTO) (*e.Passenger, error) {
	id, err := identity.ParseID(input.ID)
	if err != nil {
		return nil, err
	}

	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return nil, err
	}
	p := e.Passenger{
		ID:                id,
		Name:              input.Name,
		Nickname:          input.Nickname,
		RouteCode:         input.RouteCode,
		MonitorID:         monitorID,
		Goes:              input.Goes,
		Comesback:         input.Comesback,
		RegisterConfirmed: input.RegisterConfirmed,
	}
	if err := p.IsValid(); err != nil {
		return nil, err
	}
	return &p, nil
}

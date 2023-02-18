package factory

import (
	"sync"

	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
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

func (df *passengerFactory) Create(input PassengerInputDTO) (*entity.Passenger, error) {
	monitorID, err := identity.ParseID(input.ID)
	if err != nil {
		return nil, err
	}

	p, err := entity.NewPassenger(
		input.ID,
		input.Name,
		input.Nickname,
		input.RouteCode,
		input.Goes,
		input.Comesback,
		input.RegisterConfirmed,
		monitorID,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

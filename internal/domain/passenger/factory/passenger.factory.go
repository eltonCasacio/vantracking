package factory

import (
	"sync"

	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
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
	p, err := entity.NewPassenger(
		input.Name,
		input.Nickname,
		input.RouteCode,
		input.MonitorID,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

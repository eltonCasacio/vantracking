package repository

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PassengerRepositoryInterface interface {
	repo.RepositoryInterface[e.Passenger]
	ListNotConfirmedPassengers() ([]e.Passenger, error)
	ConfirmPassengerRegister(id string, confirm bool) error
	ListByMonitorID(monitorID string) ([]e.Passenger, error)
}

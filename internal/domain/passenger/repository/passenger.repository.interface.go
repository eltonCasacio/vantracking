package repository

import (
	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PassengerRepositoryInterface interface {
	repo.RepositoryInterface[entity.Passenger]
	ListNotConfirmedPassengers() ([]entity.Passenger, error)
	FindByNameAndNickname(name, monitorID string) (*entity.Passenger, error)
	ConfirmPassengerRegister(id string, confirm bool) error
}

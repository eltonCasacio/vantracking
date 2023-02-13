package repository

import (
	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PassengerRepositoryInterface interface {
	repo.RepositoryInterface[entity.Passenger]
}

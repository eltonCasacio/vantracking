package monitor

import repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"

type PassengerRepositoryInterface interface {
	repo.RepositoryInterface[Passenger]
}

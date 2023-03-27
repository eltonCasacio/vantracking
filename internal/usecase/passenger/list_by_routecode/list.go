package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type listAllByRouteCodeUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *listAllByRouteCodeUseCase {
	return &listAllByRouteCodeUseCase{
		repository: repository,
	}
}

func (u *listAllByRouteCodeUseCase) ListAllByRouteCode(routeCode string) ([]PassengerOutputDTO, error) {

	passengers, err := u.repository.ListByRouteCode(routeCode)
	if err != nil {
		return []PassengerOutputDTO{}, err
	}

	var passengersOutput []PassengerOutputDTO
	for _, passenger := range passengers {
		p := PassengerOutputDTO{
			ID:                passenger.ID.String(),
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			MonitorID:         passenger.MonitorID.String(),
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			RegisterConfirmed: passenger.IsRegisterConfirmed(),
			SchoolName:        passenger.SchoolName,
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

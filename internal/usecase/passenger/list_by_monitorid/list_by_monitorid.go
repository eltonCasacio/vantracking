package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type PassengersUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *PassengersUseCase {
	return &PassengersUseCase{
		repository: repository,
	}
}

func (u *PassengersUseCase) ListByMonitorID(monitorID string) ([]PassengerOutDTO, error) {
	passengers, err := u.repository.ListByMonitorID(monitorID)
	if err != nil {
		return []PassengerOutDTO{}, err
	}

	var passengersOutput []PassengerOutDTO
	for _, passenger := range passengers {
		p := PassengerOutDTO{
			ID:                passenger.ID.String(),
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			RegisterConfirmed: passenger.IsRegisterConfirmed(),
			MonitorID:         passenger.MonitorID.String(),
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

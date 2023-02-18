package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type newPassengersUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *newPassengersUseCase {
	return &newPassengersUseCase{
		repository: repository,
	}
}

func (u *newPassengersUseCase) ListNotConfirmedPassengers() ([]PassengerOutDTO, error) {
	passengers, err := u.repository.ListNotConfirmedPassengers()
	if err != nil {
		return []PassengerOutDTO{}, err
	}

	var passengersOutput []PassengerOutDTO
	for _, passenger := range passengers {
		p := PassengerOutDTO{
			ID:                passenger.GetID().String(),
			Name:              passenger.GetName(),
			Nickname:          passenger.GetNickname(),
			RouteCode:         passenger.GetRouteCode(),
			Goes:              passenger.GetGoes(),
			Comesback:         passenger.GetComesBack(),
			RegisterConfirmed: passenger.IsRegisterConfirmed(),
			MonitorID:         passenger.GetMonitorID().String(),
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

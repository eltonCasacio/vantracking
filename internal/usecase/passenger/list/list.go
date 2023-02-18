package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type listAllUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *listAllUseCase {
	return &listAllUseCase{
		repository: repository,
	}
}

func (u *listAllUseCase) ListAll() ([]PassengerOutputDTO, error) {
	passengers, err := u.repository.FindAll()
	if err != nil {
		return []PassengerOutputDTO{}, err
	}

	var passengersOutput []PassengerOutputDTO
	for _, passenger := range passengers {
		p := PassengerOutputDTO{
			ID:                passenger.GetID().String(),
			Name:              passenger.GetName(),
			Nickname:          passenger.GetNickname(),
			RouteCode:         passenger.GetRouteCode(),
			MonitorID:         passenger.GetMonitorID().String(),
			Goes:              passenger.GetGoes(),
			Comesback:         passenger.GetComesBack(),
			RegisterConfirmed: passenger.IsRegisterConfirmed(),
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

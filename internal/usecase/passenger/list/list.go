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
			ID:                passenger.ID.String(),
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			MonitorID:         passenger.MonitorID.String(),
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			RegisterConfirmed: passenger.IsRegisterConfirmed(),
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

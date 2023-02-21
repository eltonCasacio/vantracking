package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type listGoNoGoUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *listGoNoGoUseCase {
	return &listGoNoGoUseCase{
		repository: repository,
	}
}

func (u *listGoNoGoUseCase) ListGoNoGo(routeCode PassengerGoNoGoInputDTO) ([]PassengerOutputDTO, error) {
	passengers, err := u.repository.ListGoNoGoPassenger(routeCode.RouteCode)
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

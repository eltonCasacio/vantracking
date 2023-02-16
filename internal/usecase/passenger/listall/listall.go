package driver

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
			id:           passenger.GetID().String(),
			name:         passenger.GetName(),
			nickname:     passenger.GetNickname(),
			routeCode:    passenger.GetRouteCode(),
			monitorID:    passenger.GetMonitorID().String(),
			dontGo:       passenger.GetDontGo(),
			dontComeback: passenger.GetDontComeback(),
		}
		passengersOutput = append(passengersOutput, p)
	}

	return passengersOutput, nil
}

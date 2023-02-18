package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type findByIDUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *findByIDUseCase {
	return &findByIDUseCase{
		repository: repository,
	}
}

func (cd *findByIDUseCase) FindByID(id string) (PassengerOutDTO, error) {

	passenger, err := cd.repository.FindByID(id)
	if err != nil {
		return PassengerOutDTO{}, err
	}

	output := PassengerOutDTO{
		ID:                passenger.GetID().String(),
		Name:              passenger.GetName(),
		Nickname:          passenger.GetNickname(),
		RouteCode:         passenger.GetRouteCode(),
		Goes:              passenger.GetGoes(),
		Comesback:         passenger.GetComesBack(),
		RegisterConfirmed: passenger.IsRegisterConfirmed(),
		MonitorID:         passenger.GetMonitorID().String(),
	}
	return output, nil
}

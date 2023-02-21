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
		ID:                passenger.ID.String(),
		Name:              passenger.Name,
		Nickname:          passenger.Nickname,
		RouteCode:         passenger.RouteCode,
		Goes:              passenger.Goes,
		Comesback:         passenger.Comesback,
		RegisterConfirmed: passenger.IsRegisterConfirmed(),
		MonitorID:         passenger.MonitorID.String(),
	}
	return output, nil
}

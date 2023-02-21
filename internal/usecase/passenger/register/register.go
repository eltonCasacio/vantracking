package passenger

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/passenger/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type registerUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *registerUseCase {
	return &registerUseCase{
		repository: repository,
	}
}

func (u *registerUseCase) Register(input PassengerInputDTO) error {
	Input := f.NewPassengerInputDTO{
		Name:      input.Name,
		RouteCode: input.RouteCode,
		MonitorID: input.MonitorID,
	}

	passenger, err := f.PassengerFactory().NewPassenger(Input)
	if err != nil {
		return err
	}

	err = u.repository.Create(passenger)
	if err != nil {
		return err
	}

	return nil
}

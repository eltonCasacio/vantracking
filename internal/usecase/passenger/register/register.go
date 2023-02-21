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
	Input := f.PassengerInputDTO{
		ID:        "",
		Name:      input.Name,
		Nickname:  input.Nickname,
		RouteCode: input.RouteCode,
		MonitorID: input.MonitorID,
	}

	passenger, err := f.PassengerFactory().Create(Input)
	if err != nil {
		return err
	}

	err = u.repository.Create(passenger)
	if err != nil {
		return err
	}

	return nil
}

package passenger

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/passenger/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type updateUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *updateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (u *updateUseCase) Update(input PassengerOutDTO) error {
	Input := f.PassengerInputDTO{
		ID:        input.ID,
		Name:      input.Name,
		Nickname:  input.Nickname,
		RouteCode: input.RouteCode,
	}

	passenger, err := f.PassengerFactory().CreateInstance(Input)
	if err != nil {
		return err
	}

	err = u.repository.Update(passenger)
	if err != nil {
		return err
	}
	return nil
}

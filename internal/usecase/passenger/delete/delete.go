package passenger

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type deletePassengerUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *deletePassengerUseCase {
	return &deletePassengerUseCase{
		repository: repository,
	}
}

func (u *deletePassengerUseCase) Delete(id string) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

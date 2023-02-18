package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type confirmPassengerRegisterUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *confirmPassengerRegisterUseCase {
	return &confirmPassengerRegisterUseCase{
		repository: repository,
	}
}

func (u *confirmPassengerRegisterUseCase) ConfirmPassengerRegister(input PassengerInputDTO) error {
	err := u.repository.ConfirmPassengerRegister(input.PassengerID, input.RegisterConfirmed)
	if err != nil {
		return err
	}
	return nil
}

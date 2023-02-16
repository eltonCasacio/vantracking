package driver

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/monitor/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type updateUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *updateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (u *updateUseCase) Update(input DriverInputDTO) error {
	Input := f.CreateMonitorInputDTO{
		ID:          input.ID,
		Name:        input.Name,
		CPF:         input.CPF,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
	}

	monitorInstance, err := f.MonitorFactory().Create(Input)
	if err != nil {
		return err
	}

	err = u.repository.Update(monitorInstance)
	if err != nil {
		return err
	}
	return nil
}

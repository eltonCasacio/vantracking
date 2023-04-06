package monitor

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

func (u *updateUseCase) Update(input InputDTO) error {
	Input := f.InstanceMonitorInputDTO{
		ID:          input.ID,
		Name:        input.Name,
		CPF:         input.CPF,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
		Complement:  input.Complement,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
	}

	monitorInstance, err := f.MonitorFactory().Instance(Input)
	if err != nil {
		return err
	}

	err = u.repository.Update(monitorInstance)
	if err != nil {
		return err
	}
	return nil
}

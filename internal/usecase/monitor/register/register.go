package monitor

import (
	"errors"

	f "github.com/eltoncasacio/vantracking/internal/domain/monitor/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type RegisterUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *RegisterUseCase {
	return &RegisterUseCase{
		repository: repository,
	}
}

func (cd *RegisterUseCase) Register(input InputDTO) error {
	Input := f.NewMonitorInputDTO{
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

	found, _ := cd.repository.FindByCPF(monitorInstance.CPF)
	if found != nil {
		return errors.New("monitor already exists")
	}

	err = cd.repository.Create(monitorInstance)
	if err != nil {
		return err
	}
	return nil
}

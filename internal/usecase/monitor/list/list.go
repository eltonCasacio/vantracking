package monitor

import (
	"fmt"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type listUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *listUseCase {
	return &listUseCase{
		repository: repository,
	}
}

func (cd *listUseCase) List() ([]OutputDTO, error) {
	found, err := cd.repository.FindAll()
	if err != nil {
		return []OutputDTO{}, err
	}

	monitorsOutput := []OutputDTO{}

	for _, monitor := range found {
		addr := monitor.GetAddress()

		output := OutputDTO{
			ID:          monitor.GetID().String(),
			Name:        monitor.GetName(),
			CPF:         monitor.GetCPF(),
			PhoneNumber: monitor.GetPhoneNumber(),
			UF:          addr.GetUF(),
			City:        addr.GetCity(),
			Street:      addr.GetCity(),
			Number:      addr.GetNumber(),
			CEP:         fmt.Sprint(addr.GetCEP()),
		}
		monitorsOutput = append(monitorsOutput, output)
	}

	return monitorsOutput, nil
}

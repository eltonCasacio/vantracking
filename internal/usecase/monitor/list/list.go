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

	for _, driver := range found {
		addr := driver.GetAddress()

		output := OutputDTO{
			ID:          driver.GetID().String(),
			Name:        driver.GetName(),
			CPF:         driver.GetCPF(),
			PhoneNumber: driver.GetPhoneNumber(),
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

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
		addr := monitor.Address

		output := OutputDTO{
			ID:          monitor.ID.String(),
			Name:        monitor.Name,
			CPF:         monitor.CPF,
			PhoneNumber: monitor.PhoneNumber,
			UF:          addr.UF,
			City:        addr.City,
			Street:      addr.City,
			Number:      addr.Number,
			CEP:         fmt.Sprint(addr.CEP),
			Complement:  addr.Complement,
			Latitude:    addr.Latitude,
			Longitude:   addr.Longitude,
		}
		monitorsOutput = append(monitorsOutput, output)
	}

	return monitorsOutput, nil
}

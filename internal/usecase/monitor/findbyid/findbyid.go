package monitor

import (
	"fmt"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type findByIDUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *findByIDUseCase {
	return &findByIDUseCase{
		repository: repository,
	}
}

func (cd *findByIDUseCase) FindByID(id string) (OutputDTO, error) {

	found, err := cd.repository.FindByID(id)
	if err != nil {
		return OutputDTO{}, err
	}

	addr := found.Address

	output := OutputDTO{
		ID:          found.ID.String(),
		Name:        found.Name,
		CPF:         found.CPF,
		PhoneNumber: found.PhoneNumber,
		UF:          addr.UF,
		City:        addr.City,
		Street:      addr.City,
		Number:      addr.Number,
		CEP:         fmt.Sprint(addr.CEP),
		Complement:  addr.Complement,
		Latitude:    addr.Latitude,
		Longitude:   addr.Longitude,
	}
	return output, nil
}

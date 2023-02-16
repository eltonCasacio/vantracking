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

	addr := found.GetAddress()

	output := OutputDTO{
		ID:          found.GetID().String(),
		Name:        found.GetName(),
		CPF:         found.GetCPF(),
		PhoneNumber: found.GetPhoneNumber(),
		UF:          addr.GetUF(),
		City:        addr.GetCity(),
		Street:      addr.GetCity(),
		Number:      addr.GetNumber(),
		CEP:         fmt.Sprint(addr.GetCEP()),
	}
	return output, nil
}

package driver

import (
	"fmt"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type findDriverByIDUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *findDriverByIDUseCase {
	return &findDriverByIDUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *findDriverByIDUseCase) FindByID(id string) (DriverOutputDTO, error) {

	d, err := cd.driverRepository.FindByID(id)
	if err != nil {
		return DriverOutputDTO{}, err
	}
	addr := d.GetAddress()
	return DriverOutputDTO{
		ID:       d.GetID().String(),
		CPF:      d.GetCPF(),
		Name:     d.GetName(),
		Nickname: d.GetNickName(),
		Phone:    d.GetPhone(),
		UF:       addr.GetUF(),
		City:     addr.GetCity(),
		Street:   addr.GetCity(),
		Number:   addr.GetNumber(),
		CEP:      fmt.Sprint(addr.GetCEP()),
	}, nil
}

package driver

import (
	"fmt"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type findDriverUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func CreateDriverUseCase(driverRepository repo.DriverRepositoryInterface) *findDriverUseCase {
	return &findDriverUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *findDriverUseCase) Execute(input DriverInputDTO) (DriverOutputDTO, error) {

	d, err := cd.driverRepository.FindByID(input.ID)
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

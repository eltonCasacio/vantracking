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
		fmt.Println(err)
		return DriverOutputDTO{}, err
	}
	addr := d.Address
	output := DriverOutputDTO{
		ID:       d.ID.String(),
		CPF:      d.CPF,
		Name:     d.Name,
		Nickname: d.Nickname,
		Phone:    d.Phone,
		UF:       addr.UF,
		City:     addr.City,
		Street:   addr.City,
		Number:   addr.Number,
		CEP:      fmt.Sprint(addr.CEP),
	}
	return output, nil
}

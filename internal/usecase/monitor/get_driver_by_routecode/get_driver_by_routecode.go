package monitor

import (
	"fmt"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type getDriverByRouteCodeUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *getDriverByRouteCodeUseCase {
	return &getDriverByRouteCodeUseCase{
		repository: repository,
	}
}

func (cd *getDriverByRouteCodeUseCase) GetDriverByRouteCode(routeCode string) (DriverOutputDTO, error) {
	found, err := cd.repository.GetDriverByRouteCode(routeCode)
	if err != nil {
		return DriverOutputDTO{}, err
	}

	addr := found.Address
	output := DriverOutputDTO{
		ID:       found.ID.String(),
		CPF:      found.CPF,
		Name:     found.Name,
		Nickname: found.Nickname,
		Phone:    found.Phone,
		UF:       addr.UF,
		City:     addr.City,
		Street:   addr.Street,
		Number:   addr.Number,
		CEP:      fmt.Sprint(addr.CEP),
	}
	return output, nil
}

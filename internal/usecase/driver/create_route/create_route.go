package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	"github.com/eltoncasacio/vantracking/internal/domain/route"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type RegisterRouteUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *RegisterRouteUseCase {
	return &RegisterRouteUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *RegisterRouteUseCase) RegisterDriver(input CreateRouteInputDTO) error {
	driverID, err := identity.ParseID(input.DriverID)
	if err != nil {
		return err
	}
	route, err := route.NewRoute(driverID, input.Name)
	if err != nil {
		return err
	}
	err = cd.driverRepository.CreateRoute(route)
	if err != nil {
		return err
	}
	return nil
}

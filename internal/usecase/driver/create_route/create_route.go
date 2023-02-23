package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	"github.com/eltoncasacio/vantracking/internal/domain/route"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type registerRouteUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *registerRouteUseCase {
	return &registerRouteUseCase{
		driverRepository: driverRepository,
	}
}

func (cd *registerRouteUseCase) RegisterDriver(input CreateRouteInputDTO) error {
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

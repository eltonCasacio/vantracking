package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type RoutesUseCase struct {
	repository repo.DriverRepositoryInterface
}

func NewUseCase(repository repo.DriverRepositoryInterface) *RoutesUseCase {
	return &RoutesUseCase{
		repository: repository,
	}
}

func (cd *RoutesUseCase) Execute(driverID string) ([]RouteOutput, error) {
	routes, err := cd.repository.Routes(driverID)
	if err != nil {
		return nil, err
	}

	var routesOutput []RouteOutput
	for _, route := range routes {
		output := RouteOutput{
			Code:     route.Code,
			Name:     route.Name,
			DriverID: route.DriverID.String(),
			Started:  route.Started,
		}
		routesOutput = append(routesOutput, output)
	}

	return routesOutput, nil
}

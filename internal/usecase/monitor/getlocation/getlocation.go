package driver

import (
	"os"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type getDriverLocationUseCase struct {
	driverRepository repo.MonitorRepositoryInterface
}

func NewUseCase(driverRepository repo.MonitorRepositoryInterface) *getDriverLocationUseCase {
	return &getDriverLocationUseCase{
		driverRepository: driverRepository,
	}
}

func (u *getDriverLocationUseCase) Get(routeCode string) GetLocationOutputDTO {
	lat := os.Getenv("latitude_" + routeCode)
	long := os.Getenv("longitude_" + routeCode)

	return GetLocationOutputDTO{
		Latitude:  lat,
		Longitude: long,
	}
}

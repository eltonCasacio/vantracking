package driver

import (
	"fmt"
	"os"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
)

type SetDriverLocationUseCase struct {
	driverRepository repo.DriverRepositoryInterface
}

func NewUseCase(driverRepository repo.DriverRepositoryInterface) *SetDriverLocationUseCase {
	return &SetDriverLocationUseCase{
		driverRepository: driverRepository,
	}
}

func (u *SetDriverLocationUseCase) Set(input SetLocationInputDTO) error {
	if err := os.Setenv("latitude_"+input.RouteCode, input.Latitude); err != nil {
		return err
	}
	if err := os.Setenv("longitude_"+input.RouteCode, input.Longitude); err != nil {
		return err
	}

	fmt.Println(os.Getenv("latitude_" + input.RouteCode))
	fmt.Println(os.Getenv("longitude_" + input.RouteCode))
	return nil
}

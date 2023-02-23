package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	driver "github.com/eltoncasacio/vantracking/internal/usecase/driver/create_route"
	delete "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete"
	delete_route "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete_route"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/findbyid"
	findAll "github.com/eltoncasacio/vantracking/internal/usecase/driver/listall"
	register "github.com/eltoncasacio/vantracking/internal/usecase/driver/register"
	sendnotification "github.com/eltoncasacio/vantracking/internal/usecase/driver/sendnotification"
	setLocationusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/setlocation"
	update "github.com/eltoncasacio/vantracking/internal/usecase/driver/update"
)

type DriverUseCases struct {
	repository repo.DriverRepositoryInterface
}

func NewDriverUsecases(repository repo.DriverRepositoryInterface) *DriverUseCases {
	return &DriverUseCases{
		repository: repository,
	}
}

func (u *DriverUseCases) CreateRouteUsecase() (*driver.RegisterRouteUseCase, driver.CreateRouteInputDTO) {
	return driver.NewUseCase(u.repository), driver.CreateRouteInputDTO{}
}

func (u *DriverUseCases) DeleteDriverUsecase() *delete.DeleteDriverUseCase {
	return delete.NewUseCase(u.repository)
}

func (u *DriverUseCases) DeleteRouteUsecase() *delete_route.DeleteRouteUseCase {
	return delete_route.NewUseCase(u.repository)
}

func (u *DriverUseCases) FindDriverByIDUsecase() (*fusecase.FindDriverByIDUseCase, fusecase.DriverOutputDTO) {
	return fusecase.NewUseCase(u.repository), fusecase.DriverOutputDTO{}
}

func (u *DriverUseCases) FindAllDriverUsecase() (*findAll.FindAllDriverUseCase, findAll.DriverOutputDTO) {
	return findAll.NewUseCase(u.repository), findAll.DriverOutputDTO{}
}

func (u *DriverUseCases) RegisterDriverUsecase() (*register.RegisterDriverUseCase, register.DriverInputDTO) {
	return register.NewUseCase(u.repository), register.DriverInputDTO{}
}

func (u *DriverUseCases) SendNotificationUsecase() *sendnotification.RegisterDriverUseCase {
	return sendnotification.NewUseCase(u.repository)
}

func (u *DriverUseCases) UpdateDriverUsecase() (*update.UpdateDriverUseCase, update.DriverInputDTO) {
	return update.NewUseCase(u.repository), update.DriverInputDTO{}
}

func (u *DriverUseCases) SetDriverLocationUsecase() (*setLocationusecase.SetDriverLocationUseCase, setLocationusecase.SetLocationInputDTO) {
	return setLocationusecase.NewUseCase(u.repository), setLocationusecase.SetLocationInputDTO{}
}

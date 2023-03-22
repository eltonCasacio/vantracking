package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	authenticate "github.com/eltoncasacio/vantracking/internal/usecase/driver/authenticate"
	driver "github.com/eltoncasacio/vantracking/internal/usecase/driver/create_route"
	delete "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete"
	delete_route "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete_route"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/findbyid"
	findAll "github.com/eltoncasacio/vantracking/internal/usecase/driver/listall"
	register "github.com/eltoncasacio/vantracking/internal/usecase/driver/register"
	routes "github.com/eltoncasacio/vantracking/internal/usecase/driver/routes"
	sendnotification "github.com/eltoncasacio/vantracking/internal/usecase/driver/sendnotification"
	setLocationusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/setlocation"
	update "github.com/eltoncasacio/vantracking/internal/usecase/driver/update"
	"github.com/go-chi/jwtauth"
)

type DriverUseCases struct {
	repository    repo.DriverRepositoryInterface
	JWT           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewDriverUsecases(repository repo.DriverRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *DriverUseCases {
	return &DriverUseCases{
		repository:    repository,
		JWT:           jwt,
		JwtExpiriesIn: jwtExpiriesIn,
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

func (u *DriverUseCases) AuthenticateUsecase() *authenticate.AuthenticateUseCase {
	return authenticate.NewAuthenticateUseCase(u.repository, u.JWT, u.JwtExpiriesIn)
}

func (u *DriverUseCases) RoutesUsecase() *routes.RoutesUseCase {
	return routes.NewUseCase(u.repository)
}

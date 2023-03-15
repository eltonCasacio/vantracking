package repository

import (
	driver "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type MonitorRepositoryInterface interface {
	repo.RepositoryInterface[entity.Monitor]
	FindByCPF(cpf string) (*entity.Monitor, error)
	GetDriverByRouteCode(routeCode string) (*driver.Driver, error)
}

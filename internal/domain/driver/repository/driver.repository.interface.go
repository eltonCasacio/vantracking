package repository

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	r "github.com/eltoncasacio/vantracking/internal/domain/route"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type DriverRepositoryInterface interface {
	repo.RepositoryInterface[e.Driver]
	FindByCPF(cpf string) (*e.Driver, error)
	CreateRoute(route *r.Route) error
}

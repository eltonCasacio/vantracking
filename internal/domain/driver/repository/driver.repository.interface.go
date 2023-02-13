package repository

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type DriverRepositoryInterface interface {
	repo.RepositoryInterface[e.Driver]
}

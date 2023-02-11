package driver

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type DriverRepositoryInterface interface {
	repo.RepositoryInterface[Driver]
}

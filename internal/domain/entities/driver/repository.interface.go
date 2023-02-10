package entity

import (
	repo "github.com/eltoncasacio/vanmonit/internal/domain/shared"
)

type DriverRepositoryInterface interface {
	repo.RepositoryInterface[Driver]
}

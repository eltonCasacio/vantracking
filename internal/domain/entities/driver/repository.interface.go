package entity

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared"
)

type DriverRepositoryInterface interface {
	repo.RepositoryInterface[Driver]
}

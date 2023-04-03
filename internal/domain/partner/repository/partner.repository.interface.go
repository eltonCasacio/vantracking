package repository

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/partner/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PartnerRepositoryInterface interface {
	repo.RepositoryInterface[e.Partner]
	ListByCity(city string) ([]e.Partner, error)
}

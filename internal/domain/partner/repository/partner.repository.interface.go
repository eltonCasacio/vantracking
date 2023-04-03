package repository

import (
	entity "github.com/eltoncasacio/vantracking/internal/domain/partner/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PartnerRepositoryInterface interface {
	repo.RepositoryInterface[entity.Partner]
}

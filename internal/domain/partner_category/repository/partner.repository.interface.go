package repository

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/partner_category/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type PartnerCategoryRepositoryInterface interface {
	repo.RepositoryInterface[e.PartnerCategory]
}

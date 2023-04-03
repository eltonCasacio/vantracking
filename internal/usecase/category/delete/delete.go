package category

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
)

type deleteCategoryUseCase struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewUseCase(repository repo.PartnerCategoryRepositoryInterface) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{
		repository: repository,
	}
}

func (u *deleteCategoryUseCase) Delete(id string) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

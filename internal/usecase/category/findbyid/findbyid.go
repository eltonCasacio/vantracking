package category

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
)

type findByIDUseCase struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewUseCase(repository repo.PartnerCategoryRepositoryInterface) *findByIDUseCase {
	return &findByIDUseCase{
		repository: repository,
	}
}

func (cd *findByIDUseCase) FindByID(id string) (CategoryOutDTO, error) {
	partner, err := cd.repository.FindByID(id)
	if err != nil {
		return CategoryOutDTO{}, err
	}

	output := CategoryOutDTO{
		ID:   partner.ID.String(),
		Name: partner.Name,
	}
	return output, nil
}

package category

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
)

type listAllUseCase struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewUseCase(repository repo.PartnerCategoryRepositoryInterface) *listAllUseCase {
	return &listAllUseCase{
		repository: repository,
	}
}

func (u *listAllUseCase) ListAll() ([]CategoryOutpu, error) {
	categories, err := u.repository.FindAll()
	if err != nil {
		return []CategoryOutpu{}, err
	}

	var output []CategoryOutpu
	for _, partner := range categories {
		p := CategoryOutpu{
			ID:   partner.ID.String(),
			Name: partner.Name,
		}
		output = append(output, p)
	}

	return output, nil
}

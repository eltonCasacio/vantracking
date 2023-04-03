package category

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/partner_category/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type updateUseCase struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewUseCase(repository repo.PartnerCategoryRepositoryInterface) *updateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (u *updateUseCase) Update(input CategoryInputDTO) error {
	id, err := identity.ParseID(input.ID)
	if err != nil {
		return err
	}

	category := e.PartnerCategory{
		ID:   id,
		Name: input.Name,
	}
	if err != nil {
		return err
	}

	err = u.repository.Update(&category)
	if err != nil {
		return err
	}
	return nil
}

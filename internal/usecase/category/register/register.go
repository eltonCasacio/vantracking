package category

import (
	e "github.com/eltoncasacio/vantracking/internal/domain/partner_category/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner_category/repository"
	"github.com/google/uuid"
)

type registerUseCase struct {
	repository repo.PartnerCategoryRepositoryInterface
}

func NewUseCase(repository repo.PartnerCategoryRepositoryInterface) *registerUseCase {
	return &registerUseCase{
		repository: repository,
	}
}

func (u *registerUseCase) Register(input CategoryInput) error {
	category := e.PartnerCategory{
		ID:   uuid.Nil,
		Name: input.Name,
	}

	err := u.repository.Create(&category)
	if err != nil {
		return err
	}

	return nil
}

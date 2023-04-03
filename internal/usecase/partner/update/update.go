package partner

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/partner/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type updateUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *updateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (u *updateUseCase) Update(input PartnerInputDTO) error {
	partner, err := f.NewPartnerFactory().NewInstance(f.PartnerInput{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
		Complement:  input.Complement,
		CategoryID:  input.CategoryID,
	})
	if err != nil {
		return err
	}

	err = u.repository.Update(partner)
	if err != nil {
		return err
	}
	return nil
}

package partner

import (
	f "github.com/eltoncasacio/vantracking/internal/domain/partner/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type registerUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *registerUseCase {
	return &registerUseCase{
		repository: repository,
	}
}

func (u *registerUseCase) Register(input PartnerInput) error {
	partner, err := f.NewPartnerFactory().NewInstance(f.PartnerInput{
		ID:          "",
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

	err = u.repository.Create(partner)
	if err != nil {
		return err
	}

	return nil
}

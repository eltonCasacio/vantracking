package partner

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type findByIDUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *findByIDUseCase {
	return &findByIDUseCase{
		repository: repository,
	}
}

func (cd *findByIDUseCase) FindByID(id string) (PartnerOutDTO, error) {

	partner, err := cd.repository.FindByID(id)
	if err != nil {
		return PartnerOutDTO{}, err
	}

	output := PartnerOutDTO{
		ID:          partner.ID.String(),
		Name:        partner.Name,
		Description: partner.Description,
		Price:       partner.Price,
		PhoneNumber: partner.PhoneNumber,
		UF:          partner.Address.UF,
		City:        partner.Address.City,
		Street:      partner.Address.Street,
		Number:      partner.Address.Number,
		CEP:         partner.Address.CEP,
		Complement:  partner.Address.Complement,
		Latitude:    partner.Address.Latitude,
		Longitude:   partner.Address.Longitude,
		CategoryID:  partner.CategoryID.String(),
	}
	return output, nil
}

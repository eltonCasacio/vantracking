package partner

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type listByCityUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *listByCityUseCase {
	return &listByCityUseCase{
		repository: repository,
	}
}

func (u *listByCityUseCase) ListByCity(city string) ([]PartnerOutputDTO, error) {
	partners, err := u.repository.ListByCity(city)
	if err != nil {
		return []PartnerOutputDTO{}, err
	}

	var output []PartnerOutputDTO
	for _, partner := range partners {
		p := PartnerOutputDTO{
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
			CategoryID:  partner.CategoryID.String(),
		}
		output = append(output, p)
	}

	return output, nil
}

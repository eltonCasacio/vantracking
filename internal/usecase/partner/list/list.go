package partner

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type listAllUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *listAllUseCase {
	return &listAllUseCase{
		repository: repository,
	}
}

func (u *listAllUseCase) ListAll() ([]PartnerOutputDTO, error) {
	partners, err := u.repository.FindAll()
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
		}
		output = append(output, p)
	}

	return output, nil
}

package partner

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/partner/repository"
)

type deletePartnerUseCase struct {
	repository repo.PartnerRepositoryInterface
}

func NewUseCase(repository repo.PartnerRepositoryInterface) *deletePartnerUseCase {
	return &deletePartnerUseCase{
		repository: repository,
	}
}

func (u *deletePartnerUseCase) Delete(id string) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

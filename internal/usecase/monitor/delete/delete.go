package monitor

import (
	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type deleteMonitorUseCase struct {
	repository repo.MonitorRepositoryInterface
}

func NewUseCase(repository repo.MonitorRepositoryInterface) *deleteMonitorUseCase {
	return &deleteMonitorUseCase{
		repository: repository,
	}
}

func (u *deleteMonitorUseCase) Delete(id string) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

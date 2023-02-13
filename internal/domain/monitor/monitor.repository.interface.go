package monitor

import repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"

type MonitorRepositoryInterface interface {
	repo.RepositoryInterface[Monitor]
}

package repository

import (
	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
	repo "github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
)

type MonitorRepositoryInterface interface {
	repo.RepositoryInterface[entity.Monitor]
}

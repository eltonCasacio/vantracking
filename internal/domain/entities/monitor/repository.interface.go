package entity

import "github.com/eltoncasacio/vanmonit/internal/domain/shared"

type MonitorRepositoryInterface interface {
	shared.RepositoryInterface[Monitor]
}

package entity

import "github.com/eltoncasacio/vantracking/internal/domain/shared"

type MonitorRepositoryInterface interface {
	shared.RepositoryInterface[Monitor]
}

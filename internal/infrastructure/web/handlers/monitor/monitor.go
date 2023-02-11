package handlers

import (
	"net/http"

	entity "github.com/eltoncasacio/vantracking/internal/domain/monitor"
)

type MonitorHandler struct {
	repository entity.MonitorRepositoryInterface
}

func NewMonitorHandler(repo entity.MonitorRepositoryInterface) *MonitorHandler {
	return &MonitorHandler{repository: repo}
}

func (dh *MonitorHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

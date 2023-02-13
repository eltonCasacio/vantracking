package handlers

import (
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
)

type MonitorHandler struct {
	repository repo.MonitorRepositoryInterface
}

func NewMonitorHandler(repo repo.MonitorRepositoryInterface) *MonitorHandler {
	return &MonitorHandler{repository: repo}
}

func (dh *MonitorHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"net/http"

	entity "github.com/eltoncasacio/vanmonit/internal/domain/entities/driver"
)

type DriverHandler struct {
	repository entity.DriverRepositoryInterface
}

func NewDriverHandler(repo entity.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

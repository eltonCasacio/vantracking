package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	driverUsecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/create"
)

type DriverHandler struct {
	repository repo.DriverRepositoryInterface
}

func NewDriverHandler(repo repo.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto DriverInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inputDriver := driverUsecase.DriverInputDTO{
		CPF:      dto.CPF,
		Name:     dto.Name,
		Nickname: dto.Nickname,
		Phone:    dto.Phone,
		UF:       dto.UF,
		City:     dto.City,
		Street:   dto.Street,
		Number:   dto.Number,
		CEP:      dto.CEP,
	}

	driverUsecase := driverUsecase.NewDriverUseCase(dh.repository)
	err = driverUsecase.Execute(inputDriver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

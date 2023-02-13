package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	du "github.com/eltoncasacio/vantracking/internal/usecase/driver/create"
)

type DriverHandler struct {
	repository repo.DriverRepositoryInterface
}

func NewDriverHandler(repo repo.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := DriverInputDTO{}
	json.NewDecoder(r.Body).Decode(&data)

	inputDriver := du.DriverInputDTO{
		CPF:      data.CPF,
		Name:     data.Name,
		Nickname: data.Nickname,
		Phone:    data.Phone,
		UF:       data.UF,
		City:     data.City,
		Street:   data.Street,
		Number:   data.Number,
		CEP:      data.CEP,
	}

	driverUsecase := du.CreateDriverUseCase(dh.repository)
	err := driverUsecase.Execute(inputDriver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"net/http"

	f "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	driverUsecase "github.com/eltoncasacio/vantracking/internal/usecase/driver"
)

type DriverHandler struct {
	repository repo.DriverRepositoryInterface
}

func NewDriverHandler(repo repo.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	input := DriverInputDTO{}
	json.NewDecoder(r.Body).Decode(&input)

	driver := f.DriverInputDTO{
		CPF:      input.CPF,
		Name:     input.Name,
		Nickname: input.Nickname,
		Phone:    input.Phone,
		UF:       input.UF,
		City:     input.City,
		Street:   input.Street,
		Number:   input.Number,
		CEP:      input.CEP,
	}

	err := driverUsecase.CreateDriverUseCase(dh.repository).Execute(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"net/http"

	entity "github.com/eltoncasacio/vantracking/internal/domain/driver"
	driverUsecase "github.com/eltoncasacio/vantracking/internal/usecase/driver"
)

type DriverHandler struct {
	repository entity.DriverRepositoryInterface
}

func NewDriverHandler(repo entity.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	input := DriverInputDTO{}
	json.NewDecoder(r.Body).Decode(&input)

	driver := entity.DriverInputDTO{
		CPF:           input.CPF,
		Name:          input.Name,
		Nickname:      input.Nickname,
		Phone:         input.Phone,
		PlateNumber:   input.PlateNumber,
		SchoolName:    input.SchoolName,
		UFAddress:     input.UFAddress,
		CityAddress:   input.CityAddress,
		StreetAddress: input.StreetAddress,
		NumberAddress: input.NumberAddress,
		CEPAddress:    input.CEPAddress,
		UFSchool:      input.UFSchool,
		CitySchool:    input.CitySchool,
		StreetSchool:  input.StreetSchool,
		NumberSchool:  input.NumberSchool,
		CEPSchool:     input.CEPSchool,
	}

	err := driverUsecase.CreateDriverUseCase(dh.repository).Execute(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

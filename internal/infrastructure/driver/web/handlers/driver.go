package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	driverUsecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/create"
	"github.com/go-chi/chi"
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

func (dh *DriverHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	driversFound, err := dh.repository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var drivers []DriverOutputDTO
	for _, driver := range driversFound {
		addr := driver.GetAddress()
		output := DriverOutputDTO{
			ID:       driver.GetID().String(),
			CPF:      driver.GetCPF(),
			Name:     driver.GetName(),
			Nickname: driver.GetNickName(),
			Phone:    driver.GetPhone(),
			UF:       addr.GetUF(),
			City:     addr.GetCity(),
			Street:   addr.GetStreet(),
			Number:   addr.GetNumber(),
			CEP:      addr.GetCEP(),
		}
		drivers = append(drivers, output)
	}

	json.NewEncoder(w).Encode(drivers)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Consult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found, err := dh.repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	addr := found.GetAddress()
	driver := DriverOutputDTO{
		ID:       found.GetID().String(),
		CPF:      found.GetCPF(),
		Name:     found.GetName(),
		Nickname: found.GetNickName(),
		Phone:    found.GetPhone(),
		UF:       addr.GetUF(),
		City:     addr.GetCity(),
		Street:   addr.GetStreet(),
		Number:   addr.GetNumber(),
		CEP:      addr.GetCEP(),
	}

	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

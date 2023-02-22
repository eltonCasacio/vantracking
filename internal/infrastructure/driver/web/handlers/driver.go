package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	dusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/findbyid"
	fausecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/listall"
	rusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/register"
	setLocationusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/setlocation"
	upusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/update"
	"github.com/go-chi/chi"
)

type DriverHandler struct {
	repository repo.DriverRepositoryInterface
}

func NewDriverHandler(repo repo.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{repository: repo}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input rusecase.DriverInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rusecase.NewUseCase(dh.repository).RegisterDriver(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := fausecase.NewUseCase(dh.repository).ListAll()

	output := []fausecase.DriverOutputDTO{}
	for _, driver := range usecaseOutput {
		d := fausecase.DriverOutputDTO{
			ID:       driver.ID,
			CPF:      driver.CPF,
			Name:     driver.Name,
			Nickname: driver.Nickname,
			Phone:    driver.Phone,
			UF:       driver.UF,
			City:     driver.City,
			Street:   driver.Street,
			Number:   driver.Number,
			CEP:      driver.CEP,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Consult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, _ := fusecase.NewUseCase(dh.repository).FindByID(id)

	driver := fusecase.DriverOutputDTO{
		ID:       output.ID,
		CPF:      output.CPF,
		Name:     output.Name,
		Nickname: output.Nickname,
		Phone:    output.Phone,
		UF:       output.UF,
		City:     output.City,
		Street:   output.Street,
		Number:   output.Number,
		CEP:      output.CEP,
	}

	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input upusecase.DriverInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = upusecase.NewUseCase(dh.repository).Update(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := dusecase.NewUseCase(dh.repository).Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) SetLocation(w http.ResponseWriter, r *http.Request) {
	var input setLocationusecase.SetLocationInputDTO
	json.NewDecoder(r.Body).Decode(&input)

	err := setLocationusecase.NewUseCase(dh.repository).Set(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	dusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/delete"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/findbyid"
	fausecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/listall"
	rusecase "github.com/eltoncasacio/vantracking/internal/usecase/driver/register"
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
	var dto DriverInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inputDriver := rusecase.DriverInputDTO{
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

	err = rusecase.NewUseCase(dh.repository).RegisterDriver(inputDriver)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	output, err := fausecase.NewUseCase(dh.repository).ListAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var drivers []DriverOutputDTO
	for _, driver := range output {
		output := DriverOutputDTO{
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
		drivers = append(drivers, output)
	}

	json.NewEncoder(w).Encode(drivers)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Consult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := fusecase.NewUseCase(dh.repository).FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	driver := DriverOutputDTO{
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
	var input UpdateDriverInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inputDriver := upusecase.DriverInputDTO{
		ID:       input.ID,
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

	err = upusecase.NewUseCase(dh.repository).Update(inputDriver)
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

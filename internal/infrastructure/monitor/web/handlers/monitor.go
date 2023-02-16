package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
	dusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/delete"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/findbyid"
	fausecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/list"
	rusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/register"
	upusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/update"
	"github.com/go-chi/chi"
)

type monitorHandler struct {
	repository repo.MonitorRepositoryInterface
}

func NewMonitorHandler(repo repo.MonitorRepositoryInterface) *monitorHandler {
	return &monitorHandler{repository: repo}
}

func (h *monitorHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input CreateInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inputDriver := rusecase.InputDTO{
		Name:        input.Name,
		CPF:         input.CPF,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
	}

	err = rusecase.NewUseCase(h.repository).Register(inputDriver)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *monitorHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := fausecase.NewUseCase(h.repository).List()

	output := []OutputDTO{}
	for _, monitor := range usecaseOutput {
		d := OutputDTO{
			ID:          monitor.ID,
			Name:        monitor.Name,
			CPF:         monitor.CPF,
			PhoneNumber: monitor.PhoneNumber,
			UF:          monitor.UF,
			City:        monitor.City,
			Street:      monitor.Street,
			Number:      monitor.Number,
			CEP:         monitor.CEP,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (dh *monitorHandler) Consult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, _ := fusecase.NewUseCase(dh.repository).FindByID(id)

	driver := OutputDTO{
		ID:          output.ID,
		CPF:         output.CPF,
		Name:        output.Name,
		PhoneNumber: output.PhoneNumber,
		UF:          output.UF,
		City:        output.City,
		Street:      output.Street,
		Number:      output.Number,
		CEP:         output.CEP,
	}

	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

func (dh *monitorHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input UpdateInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inputDriver := upusecase.DriverInputDTO{
		ID:          input.ID,
		CPF:         input.CPF,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		UF:          input.UF,
		City:        input.City,
		Street:      input.Street,
		Number:      input.Number,
		CEP:         input.CEP,
	}

	err = upusecase.NewUseCase(dh.repository).Update(inputDriver)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *monitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
	notConfirmedUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_not_confirmed_passengers"
	listAllUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/listall"
)

type passengerHandler struct {
	repository repo.PassengerRepositoryInterface
}

func NewMonitorHandler(repo repo.PassengerRepositoryInterface) *passengerHandler {
	return &passengerHandler{repository: repo}
}

func (h *passengerHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := listAllUsecase.NewUseCase(h.repository).ListAll()

	output := []OutputDTO{}
	for _, passenger := range usecaseOutput {
		d := OutputDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListNotConfirmed(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := notConfirmedUsecase.NewUseCase(h.repository).ListNotConfirmedPassengers()

	output := []OutputDTO{}
	for _, passenger := range usecaseOutput {
		d := OutputDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

// func (h *passengerHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var input CreateInputDTO
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	inputDriver := rusecase.InputDTO{
// 		Name:        input.Name,

// 	}

// 	err = rusecase.NewUseCase(h.repository).Register(inputDriver)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func (h *monitorHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
// 	usecaseOutput, _ := fausecase.NewUseCase(h.repository).List()

// 	output := []OutputDTO{}
// 	for _, monitor := range usecaseOutput {
// 		d := OutputDTO{
// 			ID:          monitor.ID,
// 			Name:        monitor.Name,
// 			CPF:         monitor.CPF,
// 			PhoneNumber: monitor.PhoneNumber,
// 			UF:          monitor.UF,
// 			City:        monitor.City,
// 			Street:      monitor.Street,
// 			Number:      monitor.Number,
// 			CEP:         monitor.CEP,
// 		}
// 		output = append(output, d)
// 	}

// 	json.NewEncoder(w).Encode(output)
// 	w.WriteHeader(http.StatusOK)
// }

// func (dh *monitorHandler) Consult(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	output, _ := fusecase.NewUseCase(dh.repository).FindByID(id)

// 	driver := OutputDTO{
// 		ID:          output.ID,
// 		CPF:         output.CPF,
// 		Name:        output.Name,
// 		PhoneNumber: output.PhoneNumber,
// 		UF:          output.UF,
// 		City:        output.City,
// 		Street:      output.Street,
// 		Number:      output.Number,
// 		CEP:         output.CEP,
// 	}

// 	json.NewEncoder(w).Encode(driver)
// 	w.WriteHeader(http.StatusOK)
// }

// func (dh *monitorHandler) Update(w http.ResponseWriter, r *http.Request) {
// 	var input UpdateInputDTO
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil || input.ID == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	inputDriver := upusecase.DriverInputDTO{
// 		ID:          input.ID,
// 		CPF:         input.CPF,
// 		Name:        input.Name,
// 		PhoneNumber: input.PhoneNumber,
// 		UF:          input.UF,
// 		City:        input.City,
// 		Street:      input.Street,
// 		Number:      input.Number,
// 		CEP:         input.CEP,
// 	}

// 	err = upusecase.NewUseCase(dh.repository).Update(inputDriver)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func (dh *monitorHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	err := dusecase.NewUseCase(dh.repository).Delete(id)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
	confirmRegisterUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/confirm_passenger_register"
	deleteUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/delete"
	finalizeRouteUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/finalize_route"
	findUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/findbyid"
	listUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list"
	listByMonitorID "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_by_monitorid"
	list_by_routecode_usecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_by_routecode"
	notConfirmedUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_not_confirmed"
	registerUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/register"
	updateUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/update"
	"github.com/go-chi/chi"
)

type passengerHandler struct {
	repository repo.PassengerRepositoryInterface
}

func NewPassengerHandler(repo repo.PassengerRepositoryInterface) *passengerHandler {
	return &passengerHandler{repository: repo}
}

func (h *passengerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerUsecase.PassengerInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecaseInput := registerUsecase.PassengerInputDTO{
		Name:       input.Name,
		Nickname:   input.Nickname,
		RouteCode:  input.RouteCode,
		SchoolName: input.SchoolName,
		MonitorID:  input.MonitorID,
	}

	err = registerUsecase.NewUseCase(h.repository).Register(usecaseInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *passengerHandler) Find(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	passenger, _ := findUsecase.NewUseCase(dh.repository).FindByID(id)

	driver := findUsecase.PassengerOutDTO{
		ID:                passenger.ID,
		Name:              passenger.Name,
		Nickname:          passenger.Nickname,
		RouteCode:         passenger.RouteCode,
		Goes:              passenger.Goes,
		Comesback:         passenger.Comesback,
		RegisterConfirmed: passenger.RegisterConfirmed,
		SchoolName:        passenger.SchoolName,
		MonitorID:         passenger.MonitorID,
	}

	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

func (dh *passengerHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input updateUsecase.PassengerInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = updateUsecase.NewUseCase(dh.repository).Update(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *passengerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := deleteUsecase.NewUseCase(dh.repository).Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := listUsecase.NewUseCase(h.repository).ListAll()

	output := []listUsecase.PassengerOutputDTO{}
	for _, passenger := range usecaseOutput {
		d := listUsecase.PassengerOutputDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
			SchoolName:        passenger.SchoolName,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListNotConfirmed(w http.ResponseWriter, r *http.Request) {
	routeCode := chi.URLParam(r, "routeCode")
	if routeCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecaseOutput, _ := notConfirmedUsecase.NewUseCase(h.repository).ListNotConfirmed(routeCode)

	output := []notConfirmedUsecase.PassengerOutDTO{}
	for _, passenger := range usecaseOutput {
		d := notConfirmedUsecase.PassengerOutDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
			SchoolName:        passenger.SchoolName,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ConfirmPassengerRegister(w http.ResponseWriter, r *http.Request) {
	var input confirmRegisterUsecase.PassengerInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	err = confirmRegisterUsecase.NewUseCase(h.repository).ConfirmPassengerRegister(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListByMonitorID(w http.ResponseWriter, r *http.Request) {
	monitor_id := chi.URLParam(r, "monitor_id")
	if monitor_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := listByMonitorID.NewUseCase(h.repository).ListByMonitorID(monitor_id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListAllByRouteCode(w http.ResponseWriter, r *http.Request) {
	routeCode := chi.URLParam(r, "routeCode")
	if routeCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usecaseOutput, _ := list_by_routecode_usecase.NewUseCase(h.repository).ListAllByRouteCode(routeCode)

	output := []listUsecase.PassengerOutputDTO{}
	for _, passenger := range usecaseOutput {
		d := listUsecase.PassengerOutputDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
			SchoolName:        passenger.SchoolName,
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) FinalizeRoute(w http.ResponseWriter, r *http.Request) {
	var input finalizeRouteUsecase.PassengerInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	finalizeRouteUsecase.NewUseCase(h.repository).FinalizeRoute(input)
	w.WriteHeader(http.StatusOK)
}

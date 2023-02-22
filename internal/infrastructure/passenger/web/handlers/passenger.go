package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
	confirmRegisterUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/confirm_passenger_register"
	deleteUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/delete"
	findUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/findbyid"
	listUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list"
	gonogoUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_gonogo"
	notConfirmedUsecase "github.com/eltoncasacio/vantracking/internal/usecase/passenger/list_not_confirmed_passengers"
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
		Name:      input.Name,
		Nickname:  input.Nickname,
		RouteCode: input.RouteCode,
		MonitorID: input.MonitorID,
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
		}
		output = append(output, d)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

func (h *passengerHandler) ListNotConfirmed(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := notConfirmedUsecase.NewUseCase(h.repository).ListNotConfirmedPassengers()

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

func (h *passengerHandler) ListGoNoGoPassenger(w http.ResponseWriter, r *http.Request) {
	var input gonogoUsecase.PassengerGoNoGoInputDTO
	json.NewDecoder(r.Body).Decode(&input)

	outputUsecase, err := gonogoUsecase.NewUseCase(h.repository).ListGoNoGo(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var output []gonogoUsecase.PassengerOutputDTO
	for _, passenger := range outputUsecase {
		p := gonogoUsecase.PassengerOutputDTO{
			ID:                passenger.ID,
			Name:              passenger.Name,
			Nickname:          passenger.Nickname,
			RouteCode:         passenger.RouteCode,
			Goes:              passenger.Goes,
			Comesback:         passenger.Comesback,
			MonitorID:         passenger.MonitorID,
			RegisterConfirmed: passenger.RegisterConfirmed,
		}
		output = append(output, p)
	}

	json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
}

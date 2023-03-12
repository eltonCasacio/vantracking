package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
	authenticateusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/authenticate"
	dusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/delete"
	fusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/findbyid"
	getlocationusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/getlocation"
	fausecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/list"
	rusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/register"
	upusecase "github.com/eltoncasacio/vantracking/internal/usecase/monitor/update"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type monitorHandler struct {
	repository    repo.MonitorRepositoryInterface
	JWT           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewMonitorHandler(repo repo.MonitorRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *monitorHandler {
	return &monitorHandler{
		repository:    repo,
		JWT:           jwt,
		JwtExpiriesIn: jwtExpiriesIn,
	}
}

func (h *monitorHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input rusecase.InputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rusecase.NewUseCase(h.repository).Register(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *monitorHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	usecaseOutput, _ := fausecase.NewUseCase(h.repository).List()

	output := []fausecase.OutputDTO{}
	for _, monitor := range usecaseOutput {
		d := fausecase.OutputDTO{
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

	driver := fusecase.OutputDTO{
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
	var input upusecase.InputDTO
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

func (dh *monitorHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	routeCode := chi.URLParam(r, "routecode")
	if routeCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output := getlocationusecase.NewUseCase(dh.repository).Get(routeCode)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (dh *monitorHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := authenticateusecase.NewUseCase(dh.repository, dh.JWT, dh.JwtExpiriesIn).Authenticate(cpf)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

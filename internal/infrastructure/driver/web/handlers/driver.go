package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	usecases "github.com/eltoncasacio/vantracking/internal/usecase/driver"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type DriverHandler struct {
	usecases      usecases.DriverUseCases
	JWT           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewDriverHandler(repo repo.DriverRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *DriverHandler {
	return &DriverHandler{usecases: *usecases.NewDriverUsecases(repo, jwt, jwtExpiriesIn)}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	usecase, input := dh.usecases.RegisterDriverUsecase()
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = usecase.RegisterDriver(&input)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) ConsultAll(w http.ResponseWriter, r *http.Request) {
	usecase, output := dh.usecases.FindAllDriverUsecase()
	usecaseOutput, _ := usecase.ListAll()

	var outputs []interface{}

	for _, driver := range usecaseOutput {
		output.ID = driver.ID
		output.CPF = driver.CPF
		output.Name = driver.Name
		output.Nickname = driver.Nickname
		output.Phone = driver.Phone
		output.UF = driver.UF
		output.City = driver.City
		output.Street = driver.Street
		output.Number = driver.Number
		output.CEP = driver.CEP
		output.Latitude = driver.Latitude
		output.Longitude = driver.Longitude

		outputs = append(outputs, output)
	}

	json.NewEncoder(w).Encode(outputs)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Consult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	usecase, output := dh.usecases.FindDriverByIDUsecase()
	driver, err := usecase.FindByID(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	output.ID = driver.ID
	output.CPF = driver.CPF
	output.Name = driver.Name
	output.Nickname = driver.Nickname
	output.Phone = driver.Phone
	output.UF = driver.UF
	output.City = driver.City
	output.Street = driver.Street
	output.Number = driver.Number
	output.CEP = driver.CEP
	output.Latitude = driver.Latitude
	output.Longitude = driver.Longitude

	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Update(w http.ResponseWriter, r *http.Request) {
	usecase, input := dh.usecases.UpdateDriverUsecase()

	err := json.NewDecoder(r.Body).Decode(&input)

	err = usecase.Update(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	usecase := dh.usecases.DeleteDriverUsecase()

	err := usecase.Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) SetLocation(w http.ResponseWriter, r *http.Request) {
	usecase, input := dh.usecases.SetDriverLocationUsecase()
	json.NewDecoder(r.Body).Decode(&input)

	err := usecase.Set(input)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) CreateRoute(w http.ResponseWriter, r *http.Request) {
	usecase, inputDTO := dh.usecases.CreateRouteUsecase()

	input := inputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = usecase.RegisterDriver(input)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) DeleteRoute(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	usecase := dh.usecases.DeleteRouteUsecase()
	err := usecase.DeleleRoute(code)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DriverHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")
	if cpf == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecase := dh.usecases.AuthenticateUsecase()

	user, err := usecase.Authenticate(cpf)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (dh *DriverHandler) Routes(w http.ResponseWriter, r *http.Request) {
	driverID := chi.URLParam(r, "driverid")
	if driverID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecase := dh.usecases.RoutesUsecase()

	routes, err := usecase.Execute(driverID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

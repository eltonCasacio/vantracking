package handlers

import (
	"encoding/json"
	"net/http"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	usecases "github.com/eltoncasacio/vantracking/internal/usecase/driver"
	"github.com/go-chi/chi"
)

type DriverHandler struct {
	usecases usecases.DriverUseCases
}

func NewDriverHandler(repo repo.DriverRepositoryInterface) *DriverHandler {
	return &DriverHandler{usecases: *usecases.NewDriverUsecases(repo)}
}

func (dh *DriverHandler) Register(w http.ResponseWriter, r *http.Request) {
	usecase, input := dh.usecases.RegisterDriverUsecase()

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
	id := chi.URLParam(r, "id")

	usecase := dh.usecases.DeleteRouteUsecase()
	err := usecase.DeleleRoute(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

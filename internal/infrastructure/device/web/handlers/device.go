package handlers

import (
	"encoding/json"
	"net/http"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	registerUsecase "github.com/eltoncasacio/vantracking/internal/usecase/device/register"
	sendNotificationUsecase "github.com/eltoncasacio/vantracking/internal/usecase/device/send_notification"
)

type DeviceHandler struct {
	repository vo.DeviceRepositoryInterface
}

func NewDeviceHandler(repo vo.DeviceRepositoryInterface) *DeviceHandler {
	return &DeviceHandler{
		repository: repo,
	}
}

func (dh *DeviceHandler) Register(w http.ResponseWriter, r *http.Request) {
	input := registerUsecase.DeviceInputDTO{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usecase := registerUsecase.NewUseCase(dh.repository)
	err = usecase.Register(&input)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (dh *DeviceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// usecase, output := dh.usecases.FindAllDriverUsecase()
	// usecaseOutput, _ := usecase.ListAll()

	// var outputs []interface{}

	// for _, driver := range usecaseOutput {
	// 	output.ID = driver.ID
	// 	output.CPF = driver.CPF
	// 	output.Name = driver.Name
	// 	output.Nickname = driver.Nickname
	// 	output.Phone = driver.Phone
	// 	output.UF = driver.UF
	// 	output.City = driver.City
	// 	output.Street = driver.Street
	// 	output.Number = driver.Number
	// 	output.CEP = driver.CEP
	// 	output.Latitude = driver.Latitude
	// 	output.Longitude = driver.Longitude

	// 	outputs = append(outputs, output)
	// }

	// json.NewEncoder(w).Encode(outputs)
	// w.WriteHeader(http.StatusOK)
}

func (dh *DeviceHandler) Get(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	// usecase, output := dh.usecases.FindDriverByIDUsecase()
	// driver, err := usecase.FindByID(id)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// output.ID = driver.ID
	// output.CPF = driver.CPF
	// output.Name = driver.Name
	// output.Nickname = driver.Nickname
	// output.Phone = driver.Phone
	// output.UF = driver.UF
	// output.City = driver.City
	// output.Street = driver.Street
	// output.Number = driver.Number
	// output.CEP = driver.CEP
	// output.Latitude = driver.Latitude
	// output.Longitude = driver.Longitude

	// json.NewEncoder(w).Encode(driver)
	// w.WriteHeader(http.StatusOK)
}

func (dh *DeviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	// usecase, input := dh.usecases.UpdateDriverUsecase()

	// err := json.NewDecoder(r.Body).Decode(&input)

	// err = usecase.Update(input)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
}

func (dh *DeviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")

	// usecase := dh.usecases.DeleteDriverUsecase()

	// err := usecase.Delete(id)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
}

func (dh *DeviceHandler) SendPushNotification(w http.ResponseWriter, r *http.Request) {
	input := sendNotificationUsecase.DeviceInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usecase := sendNotificationUsecase.NewUseCase(dh.repository)
	output, err := usecase.SendNotification(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(output))
	w.WriteHeader(http.StatusOK)
}

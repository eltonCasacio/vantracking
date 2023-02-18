package driver

type PassengerInputDTO struct {
	PassengerID       string `json:"passenger_id"`
	RegisterConfirmed bool   `json:"register_confirmed"`
}

package handlers

type CreateInputDTO struct {
	Name              string `json:"id"`
	Nickname          string `json:"id"`
	RouteCode         string `json:"id"`
	Goes              bool   `json:"id"`
	Comesback         bool   `json:"id"`
	MonitorID         string `json:"id"`
	RegisterConfirmed bool   `json:"id"`
}

type OutputDTO struct {
	ID                string `json:"id"`
	Name              string `json:"id"`
	Nickname          string `json:"id"`
	RouteCode         string `json:"id"`
	Goes              bool   `json:"id"`
	Comesback         bool   `json:"id"`
	MonitorID         string `json:"id"`
	RegisterConfirmed bool   `json:"id"`
}

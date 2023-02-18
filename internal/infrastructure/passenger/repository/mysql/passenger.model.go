package repository

type PassengerModel struct {
	ID                string
	Name              string
	Nickname          string
	RouteCode         string
	Goes              bool
	Comesback         bool
	MonitorID         string
	active            bool
	RegisterConfirmed bool
}

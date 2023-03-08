package repository

type PassengerModel struct {
	ID                string
	Name              string
	Nickname          string
	RouteCode         string
	Goes              bool
	Comesback         bool
	active            bool
	RegisterConfirmed bool
	SchoolName        string
	MonitorID         string
}

package passenger

type PassengerOutputDTO struct {
	ID                string
	Name              string
	Nickname          string
	RouteCode         string
	MonitorID         string
	Goes              bool
	Comesback         bool
	RegisterConfirmed bool
}

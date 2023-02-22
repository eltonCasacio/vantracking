package passenger

type NewPassengerInputDTO struct {
	Name      string
	Nickname  string
	RouteCode string
	MonitorID string
}

type PassengerInputDTO struct {
	ID                string
	Name              string
	Nickname          string
	RouteCode         string
	Goes              bool
	Comesback         bool
	RegisterConfirmed bool
	MonitorID         string
}

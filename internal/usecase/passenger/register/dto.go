package passenger

type PassengerInputDTO struct {
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	RouteCode  string `json:"routeCode"`
	MonitorID  string `json:"monitorID"`
	SchoolName string `json:"schoolName"`
}

package passenger

type PassengerInputDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	RouteCode         string `json:"routeCode"`
	Goes              bool   `json:"goes"`
	Comesback         bool   `json:"comesback"`
	RegisterConfirmed bool   `json:"registerConfirmed"`
	SchoolName        string `json:"schoolName"`
	MonitorID         string `json:"monitorID"`
}

type PassengerOutDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	RouteCode         string `json:"routeCode"`
	Goes              bool   `json:"goes"`
	Comesback         bool   `json:"comesback"`
	RegisterConfirmed bool   `json:"registerConfirmed"`
	SchoolName        string `json:"schoolName"`
	MonitorID         string `json:"monitorID"`
}

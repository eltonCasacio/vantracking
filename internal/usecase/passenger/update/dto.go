package passenger

type PassengerInputDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	RouteCode         string `json:"route_code"`
	Goes              bool   `json:"goes"`
	Comesback         bool   `json:"comesback"`
	RegisterConfirmed bool   `json:"register_confirmed"`
	MonitorID         string `json:"monitor_id"`
}

type PassengerOutDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	RouteCode         string `json:"route_code"`
	Goes              bool   `json:"goes"`
	Comesback         bool   `json:"comesback"`
	RegisterConfirmed bool   `json:"register_confirmed"`
	MonitorID         string `json:"monitor_id"`
}

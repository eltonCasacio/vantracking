package passenger

type PassengerOutDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	RouteCode string `json:"route_code"`
	MonitorID string `json:"monitor_id"`
}

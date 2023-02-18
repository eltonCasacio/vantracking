package handlers

type CreateInputDTO struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	RouteCode string `json:"route_code"`
	MonitorID string `json:"monitor_id"`
}

type UpdateInputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	RouteCode string `json:"route_code"`
}

type OutputDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	RouteCode         string `json:"route_code"`
	Goes              bool   `json:"goes"`
	Comesback         bool   `json:"comesback"`
	MonitorID         string `json:"monitor_id"`
	RegisterConfirmed bool   `json:"register_confirmed"`
}

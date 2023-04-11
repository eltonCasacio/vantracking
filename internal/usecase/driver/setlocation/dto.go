package driver

type SetLocationInputDTO struct {
	RouteCode string `json:"route_code"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

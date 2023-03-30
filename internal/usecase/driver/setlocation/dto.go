package driver

type SetLocationInputDTO struct {
	RouteCode string `json:"routeCode"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

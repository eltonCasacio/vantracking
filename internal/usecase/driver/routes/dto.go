package driver

type RouteOutput struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	DriverID string `json:"driverid"`
	Started  bool   `json:"started"`
}

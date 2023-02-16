package factory

import "github.com/eltoncasacio/vantracking/pkg/identity"

type PassengerInputDTO struct {
	ID           identity.ID
	MonitorID    identity.ID
	Name         string
	Nickname     string
	RouteCode    string
	DontGo       bool
	DontComeback bool
}

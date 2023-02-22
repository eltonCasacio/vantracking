package passenger

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Passenger struct {
	ID                identity.ID
	Name              string
	Nickname          string
	RouteCode         string
	MonitorID         identity.ID
	Goes              bool
	Comesback         bool
	RegisterConfirmed bool
}

func NewPassenger(name, routeCode, nickname string, monitorID identity.ID) (*Passenger, error) {
	p := &Passenger{
		ID:                identity.NewID(),
		Name:              name,
		Nickname:          nickname,
		RouteCode:         routeCode,
		MonitorID:         monitorID,
		Goes:              true,
		Comesback:         true,
		RegisterConfirmed: false,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Passenger) IsValid() error {
	if p.Name == "" {
		return errors.New("invalid name")
	}
	if p.RouteCode == "" {
		return errors.New("invalid route code")
	}
	return nil
}

func (p *Passenger) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	p.Name = name
	return nil
}

func (p *Passenger) SetNickname(nickname string) {
	p.Nickname = nickname
}

func (p *Passenger) SetRouteCode(routeCode string) error {
	if err := routeCode == ""; err {
		return errors.New("invalid route code")
	}
	p.RouteCode = routeCode
	return nil
}

func (p *Passenger) ChangeGoNoGo(goes, comesback bool) {
	p.Goes = goes
	p.Comesback = comesback
}

func (p *Passenger) IsRegisterConfirmed() bool {
	return p.RegisterConfirmed
}

func (p *Passenger) ConfirmRegister(confirmed bool) {
	p.RegisterConfirmed = confirmed
}

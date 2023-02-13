package entity

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

var erros []error

type Passenger struct {
	id        identity.ID
	name      string
	nickname  string
	routeCode string
	monitorID identity.ID
}

func NewPassenger(name, nickname, routeCode string, monitorID identity.ID) (*Passenger, error) {
	p := &Passenger{
		id:        identity.NewID(),
		name:      name,
		nickname:  nickname,
		routeCode: routeCode,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Passenger) IsValid() error {
	if err := p.name == ""; err {
		return errors.New("name is required")
	}
	if err := p.routeCode == ""; err {
		return errors.New("route code is required")
	}
	return nil
}

func (p *Passenger) GetName() string {
	return p.name
}

func (p *Passenger) GetNickname() string {
	return p.nickname
}

func (p *Passenger) GetRouteCode() string {
	return p.routeCode
}

func (p *Passenger) GetMonitorID() identity.ID {
	return p.monitorID
}

func (p *Passenger) SetNickname(nickname string) {
	p.nickname = nickname
}

func (p *Passenger) SetRouteCode(routeCode string) error {
	if err := routeCode == ""; err {
		return errors.New("route code must be a valid route code")
	}
	p.routeCode = routeCode
	return nil
}

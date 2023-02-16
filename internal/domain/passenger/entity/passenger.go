package entity

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Passenger struct {
	id           identity.ID
	name         string
	nickname     string
	routeCode    string
	monitorID    identity.ID
	dontGo       bool
	dontComeback bool
}

func NewPassenger(name, nickname, routeCode string, monitorID identity.ID) (*Passenger, error) {
	p := &Passenger{
		id:           identity.NewID(),
		name:         name,
		nickname:     nickname,
		routeCode:    routeCode,
		monitorID:    monitorID,
		dontGo:       false,
		dontComeback: false,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Passenger) IsValid() error {
	if err := p.GetName() == ""; err {
		return errors.New("invalid name")
	}
	if err := p.GetRouteCode() == ""; err {
		return errors.New("invalid route code")
	}
	return nil
}

func (p *Passenger) GetName() string {
	return p.name
}

func (p *Passenger) GetID() identity.ID {
	return p.id
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

func (p *Passenger) GetDontGo() bool {
	return p.dontGo
}

func (p *Passenger) GetDontComeback() bool {
	return p.dontComeback
}

func (p *Passenger) SetNickname(nickname string) {
	p.nickname = nickname
}

func (p *Passenger) SetRouteCode(routeCode string) error {
	if err := routeCode == ""; err {
		return errors.New("invalid route code")
	}
	p.routeCode = routeCode
	return nil
}

func (p *Passenger) ChangeGoNoGo(dontGo, dontComeback bool) {
	p.dontGo = dontGo
	p.dontComeback = dontComeback
}

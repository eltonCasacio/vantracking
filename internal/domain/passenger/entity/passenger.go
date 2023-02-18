package entity

import (
	"errors"

	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Passenger struct {
	id                identity.ID
	name              string
	nickname          string
	routeCode         string
	monitorID         identity.ID
	goes              bool
	comesback         bool
	registerConfirmed bool
}

func NewPassenger(id, name, nickname, routeCode string, goes, comesback, registerConfirmed bool, monitorID identity.ID) (*Passenger, error) {
	newID, err := identity.ParseID(id)
	if err != nil {
		newID = identity.NewID()
	}

	p := &Passenger{
		id:                newID,
		name:              name,
		nickname:          nickname,
		routeCode:         routeCode,
		monitorID:         monitorID,
		goes:              goes,
		comesback:         comesback,
		registerConfirmed: registerConfirmed,
	}

	err = p.IsValid()
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
	if err := p.GetMonitorID().String() == ""; err {
		return errors.New("invalid monitor id")
	}
	_, err := identity.ParseID(p.GetMonitorID().String())
	if err != nil {
		return errors.New("invalid monitor id")
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

func (p *Passenger) GetGoes() bool {
	return p.goes
}

func (p *Passenger) GetComesBack() bool {
	return p.comesback
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

func (p *Passenger) ChangeGoNoGo(goes, comesback bool) {
	p.goes = goes
	p.comesback = comesback
}

func (p *Passenger) IsRegisterConfirmed() bool {
	return p.registerConfirmed
}

func (p *Passenger) ConfirmeRegister(confirmed bool) {
	p.registerConfirmed = confirmed
}

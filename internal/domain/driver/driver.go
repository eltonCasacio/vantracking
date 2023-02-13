package driver

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/entity"
)

type Driver struct {
	id       entity.ID
	cpf      string
	name     string
	nickname string
	address  vo.Address
	routes   []Route
}

func newDriver(cpf, name, nickname, phone string, address vo.Address) (*Driver, error) {
	d := &Driver{
		id:       entity.NewID(),
		cpf:      cpf,
		name:     name,
		nickname: nickname,
		address:  address,
	}

	err := d.IsValid()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Driver) IsValid() error {
	if err := d.cpf == ""; err {
		return errors.New("cpf is required")
	}
	if err := d.name == ""; err {
		return errors.New("name is required")
	}
	if err := d.address.IsValid(); err != nil {
		return errors.New("address is invalid")
	}
	return nil
}

func (d *Driver) GetID() entity.ID {
	return d.id
}

func (d *Driver) GetCPF() string {
	return d.cpf
}

func (d *Driver) GetName() string {
	return d.name
}

func (d *Driver) GetNickName() string {
	return d.nickname
}

func (d *Driver) GetAddress() vo.Address {
	return d.address
}

func (d *Driver) AddRoute(route Route) error {
	if err := route.IsValid(); err != nil {
		return err
	}
	d.routes = append(d.routes, route)
	return nil
}

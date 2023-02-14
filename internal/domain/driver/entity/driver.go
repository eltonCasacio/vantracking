package entity

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Driver struct {
	id       identity.ID
	cpf      string
	name     string
	nickname string
	phone    string
	address  vo.Address
}

func NewDriver(cpf, name, nickname, phone string, address vo.Address) (*Driver, error) {
	d := &Driver{
		id:       identity.NewID(),
		cpf:      cpf,
		name:     name,
		nickname: nickname,
		phone:    phone,
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
		return errors.New("cpf invalid")
	}
	if err := d.name == ""; err {
		return errors.New("name invalid")
	}
	if err := d.address.IsValid(); err != nil {
		return err
	}
	if err := d.GetPhone() == ""; err {
		return errors.New("phone invalid")
	}
	return nil
}

func (d *Driver) GetID() identity.ID {
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

func (d *Driver) GetPhone() string {
	return d.phone
}

func (d *Driver) GetAddress() vo.Address {
	return d.address
}

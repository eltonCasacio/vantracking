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

func NewDriver(id, cpf, name, nickname, phone string, address vo.Address) (*Driver, error) {
	newID, err := identity.ParseID(id)
	if err != nil {
		newID = identity.NewID()
	}

	d := &Driver{
		id:       newID,
		cpf:      cpf,
		name:     name,
		nickname: nickname,
		phone:    phone,
		address:  address,
	}

	err = d.IsValid()
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

func (d *Driver) ChangeName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	d.name = name
	return nil
}

func (d *Driver) ChangeNickname(nickname string) {
	d.nickname = nickname
}

func (d *Driver) ChangeCPF(cpf string) error {
	if cpf == "" {
		return errors.New("invalid cpf")
	}
	d.cpf = cpf
	return nil
}

func (d *Driver) ChangePhone(phone string) error {
	if phone == "" {
		return errors.New("invalid phone")
	}
	d.phone = phone
	return nil
}

func (d *Driver) ChangeAddress(address vo.Address) error {
	if err := address.IsValid(); err != nil {
		return errors.New("invalid address")
	}
	d.address = address
	return nil
}

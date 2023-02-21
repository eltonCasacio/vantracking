package entity

import (
	"errors"

	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type Driver struct {
	ID       identity.ID
	CPF      string
	Name     string
	Nickname string
	Phone    string
	Address  vo.Address
}

func NewDriver(cpf, name, phone string, address vo.Address) (*Driver, error) {
	d := &Driver{
		ID:      identity.NewID(),
		CPF:     cpf,
		Name:    name,
		Phone:   phone,
		Address: address,
	}

	err := d.IsValid()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Driver) IsValid() error {
	if err := d.CPF == ""; err {
		return errors.New("cpf invalid")
	}
	if err := d.Name == ""; err {
		return errors.New("name invalid")
	}
	if err := d.Address.IsValid(); err != nil {
		return err
	}
	if err := d.Phone == ""; err {
		return errors.New("phone invalid")
	}
	return nil
}

func (d *Driver) ChangeName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	d.Name = name
	return nil
}

func (d *Driver) ChangeNickname(nickname string) {
	d.Nickname = nickname
}

func (d *Driver) ChangeCPF(cpf string) error {
	if cpf == "" {
		return errors.New("invalid cpf")
	}
	d.CPF = cpf
	return nil
}

func (d *Driver) ChangePhone(phone string) error {
	if phone == "" {
		return errors.New("invalid phone")
	}
	d.Phone = phone
	return nil
}

func (d *Driver) ChangeAddress(address vo.Address) error {
	if err := address.IsValid(); err != nil {
		return errors.New("invalid address")
	}
	d.Address = address
	return nil
}

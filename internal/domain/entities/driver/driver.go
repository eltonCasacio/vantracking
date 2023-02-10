package entity

import (
	a "github.com/eltoncasacio/vanmonit/internal/domain/value_objects"
	"github.com/eltoncasacio/vanmonit/pkg/entity"
)

type Driver struct {
	id           entity.ID
	cpf          string
	name         string
	nickname     string
	phoneNumber  string
	plate_number string
	schoolCode   []string
	address      a.Address
	code         string
}

func NewDriver(cpf, name, nickname, phone, plateNumber string, schoolCode []string, address *a.Address) (*Driver, error) {
	d := &Driver{
		id:           entity.NewID(),
		cpf:          cpf,
		name:         name,
		nickname:     nickname,
		phoneNumber:  phone,
		plate_number: plateNumber,
		schoolCode:   schoolCode,
		address:      *address,
	}

	err := d.IsValid()
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Driver) IsValid() error {
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

func (d *Driver) GetPhoneNumber() string {
	return d.phoneNumber
}

func (d *Driver) GetPlateNumber() string {
	return d.plate_number
}

func (d *Driver) GetAddress() a.Address {
	return d.address
}

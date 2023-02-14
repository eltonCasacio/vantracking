package valueobjects

import "errors"

type Address struct {
	uf     string
	city   string
	street string
	number string
	cep    string
}

func NewAddress(uf, city, street, number, cep string) (*Address, error) {
	addr := &Address{
		uf:     uf,
		city:   city,
		street: street,
		number: number,
		cep:    cep,
	}
	err := addr.IsValid()
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *Address) IsValid() error {
	if a.GetUF() == "" {
		return errors.New("uf invalid")
	}
	if a.GetCity() == "" {
		return errors.New("city invalid")
	}
	if a.GetStreet() == "" {
		return errors.New("street invalid")
	}
	if a.GetCEP() == "" {
		return errors.New("invalid cep")
	}

	return nil
}

func (a *Address) GetUF() string {
	return a.uf
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) GetStreet() string {
	return a.street
}

func (a *Address) GetNumber() string {
	return a.number
}

func (a *Address) GetCEP() string {
	return a.cep
}

package valueobjects

import "errors"

type Address struct {
	UF     string
	City   string
	Street string
	Number string
	CEP    string
}

func NewAddress(uf, city, street, number, cep string) (*Address, error) {
	addr := &Address{
		UF:     uf,
		City:   city,
		Street: street,
		Number: number,
		CEP:    cep,
	}
	err := addr.IsValid()
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *Address) IsValid() error {
	if a.UF == "" {
		return errors.New("uf invalid")
	}
	if a.City == "" {
		return errors.New("city invalid")
	}
	if a.Street == "" {
		return errors.New("street invalid")
	}
	if a.CEP == "" {
		return errors.New("invalid cep")
	}

	return nil
}

package valueobjects

type Address struct {
	uf     string
	city   string
	street string
	number string
	cep    int
}

func NewAddresses(uf string, city string, street string, number string, cep int) (*Address, error) {
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

func (a *Address) GetCEP() int {
	return a.cep
}

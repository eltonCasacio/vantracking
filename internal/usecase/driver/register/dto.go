package driver

type DriverInputDTO struct {
	CPF      string
	Name     string
	Nickname string
	Phone    string
	UF       string
	City     string
	Street   string
	Number   string
	CEP      string
}

type DriverOutputDTO struct {
	ID string
}

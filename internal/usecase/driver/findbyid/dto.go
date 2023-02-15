package driver

type DriverInputDTO struct {
	ID string
}

type DriverOutputDTO struct {
	ID       string
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

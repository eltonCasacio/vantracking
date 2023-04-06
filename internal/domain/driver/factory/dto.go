package factory

type NewDriverInputDTO struct {
	CPF        string
	Name       string
	Nickname   string
	Phone      string
	UF         string
	City       string
	Street     string
	Number     string
	CEP        string
	Complement string
	Latitude   string
	Longitude  string
}

type CreateInstanceDriverInputDTO struct {
	ID         string
	CPF        string
	Name       string
	Nickname   string
	Phone      string
	UF         string
	City       string
	Street     string
	Number     string
	CEP        string
	Complement string
	Latitude   string
	Longitude  string
}

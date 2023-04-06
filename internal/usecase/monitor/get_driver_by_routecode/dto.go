package monitor

type DriverOutputDTO struct {
	ID         string `json:"id"`
	CPF        string `json:"cpf"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	UF         string `json:"uf"`
	City       string `json:"city"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	CEP        string `json:"cep"`
	Complement string `json:"complement"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

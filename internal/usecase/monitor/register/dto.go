package monitor

type InputDTO struct {
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	PhoneNumber string `json:"phone_number"`
	UF          string `json:"uf"`
	City        string `json:"city"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	CEP         string `json:"cep"`
	Complement  string `json:"complement"`
}

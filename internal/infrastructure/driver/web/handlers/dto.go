package handlers

type DriverInputDTO struct {
	CPF      string `json:"cpf"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	UF       string `json:"uf"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Number   string `json:"number"`
	CEP      int    `json:"cep"`
}

type DriverOutputDTO struct {
	ID       string `json:"id"`
	CPF      string `json:"cpf"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	UF       string `json:"uf"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Number   string `json:"number"`
	CEP      int    `json:"cep"`
}

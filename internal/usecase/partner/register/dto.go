package partner

type PartnerInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	UF          string  `json:"uf"`
	City        string  `json:"city"`
	Street      string  `json:"street"`
	Number      string  `json:"number"`
	CEP         string  `json:"cep"`
	Complement  string  `json:"complement"`
	Latitude    string  `json:"latitude"`
	Longitude   string  `json:"longitude"`
	CategoryID  string  `json:"category_id"`
}

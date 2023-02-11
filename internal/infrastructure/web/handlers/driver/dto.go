package handlers

type DriverInputDTO struct {
	CPF           string `json:"cpf"`
	Name          string `json:"name"`
	Nickname      string `json:"nickname"`
	Phone         string `json:"phone"`
	PlateNumber   string `json:"plate_number"`
	SchoolName    string `json:"school_name"`
	UFAddress     string `json:"uf_address"`
	CityAddress   string `json:"city_address"`
	StreetAddress string `json:"street_address"`
	NumberAddress string `json:"number_address"`
	CEPAddress    int    `json:"cep_address"`
	UFSchool      string `json:"uf_school"`
	CitySchool    string `json:"city_school"`
	StreetSchool  string `json:"street_school"`
	NumberSchool  string `json:"number_school"`
	CEPSchool     int    `json:"cep_school"`
}

type DriverOutputDTO struct {
	ID            string `json:"id"`
	CPF           string `json:"cpf"`
	Name          string `json:"name"`
	Nickname      string `json:"nickname"`
	Phone         string `json:"phone"`
	PlateNumber   string `json:"plate_number"`
	SchoolName    string `json:"school_name"`
	UFAddress     string `json:"uf_address"`
	CityAddress   string `json:"city_address"`
	StreetAddress string `json:"street_address"`
	NumberAddress string `json:"number_address"`
	CEPAddress    int    `json:"cep_address"`
	UFSchool      string `json:"uf_school"`
	CitySchool    string `json:"city_school"`
	StreetSchool  string `json:"street_school"`
	NumberSchool  string `json:"number_school"`
	CEPSchool     int    `json:"cep_school"`
}

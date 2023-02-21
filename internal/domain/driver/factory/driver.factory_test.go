package factory

import (
	"testing"

	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
)

func TestDriverFactoryCreate_ErrorInvalidInput(t *testing.T) {
	factory := DriverFactory()
	assert.NotNil(t, factory)

	driver, err := factory.New(NewDriverInputDTO{})
	assert.Nil(t, driver)
	assert.NotNil(t, err)
}

func TestDriverFactoryCreate(t *testing.T) {
	input := NewDriverInputDTO{
		CPF:    "any_cpf",
		Name:   "any_name",
		Phone:  "any_phone",
		UF:     "any_uf",
		City:   "any_city",
		Street: "any_street",
		Number: "any_number",
		CEP:    "123",
	}
	factory := DriverFactory()
	assert.NotNil(t, factory)

	d, err := factory.New(input)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.NotEmpty(t, d.ID)
	assert.Equal(t, d.CPF, "any_cpf")
	assert.Equal(t, d.Name, "any_name")
	assert.Equal(t, d.Nickname, "")
	assert.Equal(t, d.Phone, "any_phone")
	assert.NotEmpty(t, d.Address)

	addr := d.Address
	assert.Equal(t, addr.CEP, "123")
	assert.Equal(t, addr.City, "any_city")
	assert.Equal(t, addr.Number, "any_number")
}

func TestDriverFactoryCreate_Error(t *testing.T) {
	input := NewDriverInputDTO{
		CPF:    "",
		Name:   "any_name",
		Phone:  "any_phone",
		UF:     "any_uf",
		City:   "any_city",
		Street: "any_street",
		Number: "any_number",
		CEP:    "123",
	}
	factory := DriverFactory()
	assert.NotNil(t, factory)

	d, err := factory.New(input)
	assert.Nil(t, d)
	assert.NotNil(t, err)
}

func TestCreateInstance(t *testing.T) {
	input := CreateInstanceDriverInputDTO{
		ID:       identity.NewID().String(),
		CPF:      "any_cpf",
		Name:     "any_name",
		Nickname: "nickname",
		Phone:    "any_phone",
		UF:       "any_uf",
		City:     "any_city",
		Street:   "any_street",
		Number:   "any_number",
		CEP:      "123",
	}
	factory := DriverFactory()

	d, err := factory.CreateInstance(input)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.NotEmpty(t, d.ID)
	assert.Equal(t, d.CPF, "any_cpf")
	assert.Equal(t, d.Name, "any_name")
	assert.Equal(t, d.Nickname, "nickname")
	assert.Equal(t, d.Phone, "any_phone")
	assert.Equal(t, d.Address.CEP, "123")
}

func TestCreateInstance_InvalidAddress(t *testing.T) {
	input := CreateInstanceDriverInputDTO{
		ID:       identity.NewID().String(),
		CPF:      "any_cpf",
		Name:     "any_name",
		Nickname: "nickname",
		Phone:    "any_phone",
		UF:       "any_uf",
		City:     "any_city",
		Street:   "any_street",
		Number:   "any_number",
		CEP:      "",
	}
	factory := DriverFactory()

	d, err := factory.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, d)
}

func TestCreateInstance_InvalidID(t *testing.T) {
	input := CreateInstanceDriverInputDTO{
		ID:       "",
		CPF:      "any_cpf",
		Name:     "any_name",
		Nickname: "nickname",
		Phone:    "any_phone",
		UF:       "any_uf",
		City:     "any_city",
		Street:   "any_street",
		Number:   "any_number",
		CEP:      "124234",
	}
	factory := DriverFactory()

	d, err := factory.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, d)
}

func TestCreateInstance_InvalidParams(t *testing.T) {
	input := CreateInstanceDriverInputDTO{
		ID:       identity.NewID().String(),
		CPF:      "any_cpf",
		Name:     "",
		Nickname: "nickname",
		Phone:    "any_phone",
		UF:       "any_uf",
		City:     "any_city",
		Street:   "any_street",
		Number:   "any_number",
		CEP:      "124234",
	}
	factory := DriverFactory()

	d, err := factory.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, d)
}

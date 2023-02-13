package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriverFactoryCreate_ErrorInvalidInput(t *testing.T) {
	factory := DriverFactory()
	assert.NotNil(t, factory)

	driver, err := factory.Create(DriverInputDTO{})
	assert.Nil(t, driver)
	assert.NotNil(t, err)
}

func TestDriverFactoryCreate(t *testing.T) {
	input := DriverInputDTO{
		CPF:    "any_cpf",
		Name:   "any_name",
		Phone:  "any_phone",
		UF:     "any_uf",
		City:   "any_city",
		Street: "any_street",
		Number: "any_number",
		CEP:    123,
	}
	factory := DriverFactory()
	assert.NotNil(t, factory)

	d, err := factory.Create(input)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.NotEmpty(t, d.GetID())
	assert.Equal(t, d.GetCPF(), "any_cpf")
	assert.Equal(t, d.GetName(), "any_name")
	assert.Equal(t, d.GetNickName(), "")
	assert.Equal(t, d.GetPhone(), "any_phone")
	assert.NotEmpty(t, d.GetAddress())

	addr := d.GetAddress()
	assert.Equal(t, addr.GetCEP(), 123)
	assert.Equal(t, addr.GetCity(), "any_city")
	assert.Equal(t, addr.GetNumber(), "any_number")
}

func TestDriverFactoryCreate_Error(t *testing.T) {
	input := DriverInputDTO{
		CPF:    "",
		Name:   "any_name",
		Phone:  "any_phone",
		UF:     "any_uf",
		City:   "any_city",
		Street: "any_street",
		Number: "any_number",
		CEP:    123,
	}
	factory := DriverFactory()
	assert.NotNil(t, factory)

	d, err := factory.Create(input)
	assert.Nil(t, d)
	assert.NotNil(t, err)

}

package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonitorFactory_SigleInstance(t *testing.T) {
	f := MonitorFactory()
	factory := MonitorFactory()
	assert.Equal(t, factory, f)
}

func TestMonitorFactory_Create(t *testing.T) {
	factory := MonitorFactory()

	input := CreateMonitorInputDTO{
		Name:        "any_name",
		CPF:         "any_cpf",
		PhoneNumber: "any_phone",
		UF:          "any_uf",
		City:        "any_city",
		Street:      "any_street",
		Number:      "any_number",
		CEP:         "123",
	}

	m, err := factory.Create(input)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestMonitorFactory_Create_InvalidMonitorData(t *testing.T) {
	factory := MonitorFactory()

	input := CreateMonitorInputDTO{
		Name:        "",
		CPF:         "any_cpf",
		PhoneNumber: "any_phone",
		UF:          "any_uf",
		City:        "any_city",
		Street:      "any_street",
		Number:      "any_number",
		CEP:         "123",
	}

	m, err := factory.Create(input)
	assert.Nil(t, m)
	assert.NotNil(t, err)
}

func TestMonitorFactory_Create_InvalidAddressData(t *testing.T) {
	factory := MonitorFactory()

	input := CreateMonitorInputDTO{
		Name:        "",
		CPF:         "any_cpf",
		PhoneNumber: "any_phone",
		UF:          "any_uf",
		City:        "any_city",
		Street:      "any_street",
		Number:      "",
		CEP:         "",
	}

	m, err := factory.Create(input)
	assert.Nil(t, m)
	assert.NotNil(t, err)
}

package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	addr, err := NewAddress("uf", "any_city", "any_street", "any_number", "123123", "", "", "")
	assert.Nil(t, err)
	assert.NotEmpty(t, addr)
	assert.Equal(t, addr.CEP, "123123")
	assert.Equal(t, addr.UF, "uf")
	assert.Equal(t, addr.City, "any_city")
	assert.Equal(t, addr.Street, "any_street")
	assert.Equal(t, addr.Number, "any_number")
}

func TestNewAddresses_ErrorUFInvalid(t *testing.T) {
	addr, err := NewAddress("", "any_city", "any_street", "any_number", "123123", "", "", "")
	assert.NotNil(t, err)
	assert.Empty(t, addr)
	assert.EqualError(t, err, "uf invalid")
}

func TestNewAddresses_ErrorCityInvalid(t *testing.T) {
	addr, err := NewAddress("any_uf", "", "any_street", "any_number", "123123", "", "", "")
	assert.NotNil(t, err)
	assert.Empty(t, addr)
	assert.EqualError(t, err, "city invalid")
}

func TestNewAddresses_ErrorStreetInvalid(t *testing.T) {
	addr, err := NewAddress("any_uf", "any_city", "", "any_number", "123123", "", "", "")
	assert.NotNil(t, err)
	assert.Empty(t, addr)
	assert.EqualError(t, err, "street invalid")
}

func TestNewAddresses_ErrorCEPInvalid_NegativeNumber(t *testing.T) {
	addr, err := NewAddress("any_uf", "any_city", "any_street", "any_number", "", "", "", "")
	assert.NotNil(t, err)
	assert.Empty(t, addr)
	assert.EqualError(t, err, "invalid cep")
}

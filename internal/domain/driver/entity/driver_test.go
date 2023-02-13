package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestNewDriver_ErrorAddresEmpty(t *testing.T) {
	d, err := NewDriver("234534534", "any_name", "any_nickname", "24534534", valueobjects.Address{})
	assert.NotNil(t, err)
	assert.Empty(t, d)
}

func TestNewDriver_ErrorCPFInvalid(t *testing.T) {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)

	d, err := NewDriver("", "any_name", "any_nickname", "24534534", *addr)
	assert.NotNil(t, err)
	assert.Empty(t, d)
	assert.EqualError(t, err, "cpf invalid")
}

func TestNewDriver_ErrorNameInvalid(t *testing.T) {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)

	d, err := NewDriver("any_cpf", "", "any_nickname", "24534534", *addr)
	assert.NotNil(t, err)
	assert.Empty(t, d)
	assert.EqualError(t, err, "name invalid")
}

func TestNewDriver_ErrorPhoneInvalid(t *testing.T) {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)

	d, err := NewDriver("any_cpf", "any_name", "", "", *addr)
	assert.NotNil(t, err)
	assert.Empty(t, d)
	assert.EqualError(t, err, "phone invalid")
}

func TestNewDriver_WithoutNickname(t *testing.T) {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)

	d, err := NewDriver("any_cpf", "any_name", "", "234325", *addr)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.NotNil(t, d.address)
	assert.Equal(t, d.cpf, "any_cpf")
	assert.Equal(t, d.name, "any_name")
	assert.Equal(t, d.nickname, "")
	assert.Equal(t, d.phone, "234325")
}

func TestNewDriver(t *testing.T) {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)

	d, err := NewDriver("any_cpf", "any_name", "any_nickname", "234325", *addr)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.NotNil(t, d.address)
	assert.Equal(t, d.cpf, "any_cpf")
	assert.Equal(t, d.name, "any_name")
	assert.Equal(t, d.nickname, "any_nickname")
	assert.Equal(t, d.phone, "234325")
}

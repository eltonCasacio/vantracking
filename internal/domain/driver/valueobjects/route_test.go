package valueobjects

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestNewRoute(t *testing.T) {
	driverAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	destinuAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "777", 8)
	driver, _ := entity.NewDriver("any_cpf", "any_name", "any_nickname", "234325", *driverAddr)

	route, err := NewRoute(driver.GetID(), "1", "alves aranha manha", *driverAddr, *destinuAddr)
	assert.Nil(t, err)
	assert.NotNil(t, route)
	assert.Equal(t, route.GetCode(), "1")
	assert.Equal(t, route.GetName(), "alves aranha manha")
	assert.EqualValues(t, route.GetOrigin(), *driverAddr)
	assert.EqualValues(t, route.GetDestiny(), *destinuAddr)
	assert.NotEmpty(t, route.GetDriverID())
}

func TestNewRoute_InvalidCode(t *testing.T) {
	driverAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	destinuAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "777", 8)
	driver, _ := entity.NewDriver("any_cpf", "any_name", "any_nickname", "234325", *driverAddr)

	_, err := NewRoute(driver.GetID(), "", "alves aranha manha", *driverAddr, *destinuAddr)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid code")
}

func TestNewRoute_InvalidName(t *testing.T) {
	driverAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	destinuAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "777", 8)
	driver, _ := entity.NewDriver("any_cpf", "any_name", "any_nickname", "234325", *driverAddr)

	_, err := NewRoute(driver.GetID(), "1", "", *driverAddr, *destinuAddr)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid name")
}

func TestNewRoute_ErrorOrigin(t *testing.T) {
	driverAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	destinuAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "777", 8)
	driver, _ := entity.NewDriver("any_cpf", "any_name", "any_nickname", "234325", *driverAddr)

	route, err := NewRoute(driver.GetID(), "1", "alves aranha manha", valueobjects.Address{}, *destinuAddr)
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

func TestNewRoute_ErrorDestiny(t *testing.T) {
	driverAddr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	driver, _ := entity.NewDriver("any_cpf", "any_name", "any_nickname", "234325", *driverAddr)
	route, err := NewRoute(driver.GetID(), "1", "alves aranha manha", *driverAddr, valueobjects.Address{})
	assert.Nil(t, route)
	assert.NotNil(t, err)
}

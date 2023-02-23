package route

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouteTestSuite struct {
	suite.Suite
	DriverAddress  valueobjects.Address
	DestinyAddress valueobjects.Address
	Driver         entity.Driver
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RouteTestSuite))
}

func (suite *RouteTestSuite) SetupTest() {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", "123")
	addrDestiny, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "2", "77")
	driver, _ := entity.NewDriver("any_cpf", "any_name", "234325", *addr)

	suite.DriverAddress = *addr
	suite.DestinyAddress = *addrDestiny
	suite.Driver = *driver
}

func (s *RouteTestSuite) TestNewRoute() {
	route, err := NewRoute(s.Driver.ID, "alves aranha manha")
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), route)
	assert.Equal(s.T(), route.Code, "1")
	assert.Equal(s.T(), route.Name, "alves aranha manha")
	assert.NotEmpty(s.T(), route.DriverID)
}

func (s *RouteTestSuite) TestNewRoute_InvalidName() {
	_, err := NewRoute(s.Driver.ID, "")
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid name")
}

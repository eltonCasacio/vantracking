package valueobjects

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
	driver, _ := entity.NewDriver("", "any_cpf", "any_name", "any_nickname", "234325", *addr)

	suite.DriverAddress = *addr
	suite.DestinyAddress = *addrDestiny
	suite.Driver = *driver
}

func (s *RouteTestSuite) TestNewRoute() {
	route, err := NewRoute(s.Driver.GetID(), "1", "alves aranha manha", s.DriverAddress, s.DestinyAddress)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), route)
	assert.Equal(s.T(), route.GetCode(), "1")
	assert.Equal(s.T(), route.GetName(), "alves aranha manha")
	assert.EqualValues(s.T(), route.GetOrigin(), s.DriverAddress)
	assert.EqualValues(s.T(), route.GetDestiny(), s.DestinyAddress)
	assert.NotEmpty(s.T(), route.GetDriverID())
}

func (s *RouteTestSuite) TestNewRoute_InvalidCode() {
	_, err := NewRoute(s.Driver.GetID(), "", "alves aranha manha", s.DriverAddress, s.DestinyAddress)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid code")
}

func (s *RouteTestSuite) TestNewRoute_InvalidName() {
	_, err := NewRoute(s.Driver.GetID(), "1", "", s.DriverAddress, s.DestinyAddress)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid name")
}

func (s *RouteTestSuite) TestNewRoute_ErrorOrigin() {
	route, err := NewRoute(s.Driver.GetID(), "1", "alves aranha manha", valueobjects.Address{}, s.DestinyAddress)
	assert.Nil(s.T(), route)
	assert.NotNil(s.T(), err)
}

func (s *RouteTestSuite) TestNewRoute_ErrorDestiny() {
	route, err := NewRoute(s.Driver.GetID(), "1", "alves aranha manha", s.DriverAddress, valueobjects.Address{})
	assert.Nil(s.T(), route)
	assert.NotNil(s.T(), err)
}

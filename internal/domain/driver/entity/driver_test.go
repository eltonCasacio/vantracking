package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DriverTestSuite struct {
	suite.Suite
	Address valueobjects.Address
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DriverTestSuite))
}

func (suite *DriverTestSuite) SetupTest() {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", "123")
	suite.Address = *addr
}

func (s *DriverTestSuite) TestNewDriver_ErrorAddresEmpty() {
	d, err := NewDriver("", "234534534", "any_name", "any_nickname", "24534534", valueobjects.Address{})
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
}

func (s *DriverTestSuite) TestNewDriver_ErrorCPFInvalid() {
	d, err := NewDriver("", "", "any_name", "any_nickname", "24534534", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "cpf invalid")
}

func (s *DriverTestSuite) TestNewDriver_ErrorNameInvalid() {
	d, err := NewDriver("", "any_cpf", "", "any_nickname", "24534534", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "name invalid")
}

func (s *DriverTestSuite) TestNewDriver_ErrorPhoneInvalid() {
	d, err := NewDriver("", "any_cpf", "any_name", "", "", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "phone invalid")
}

func (s *DriverTestSuite) TestNewDriver_WithoutNickname() {
	d, err := NewDriver("", "any_cpf", "any_name", "", "234325", s.Address)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), d)
	assert.NotNil(s.T(), d.address)
	assert.Equal(s.T(), d.cpf, "any_cpf")
	assert.Equal(s.T(), d.name, "any_name")
	assert.Equal(s.T(), d.nickname, "")
	assert.Equal(s.T(), d.phone, "234325")
}

func (s *DriverTestSuite) TestNewDriver() {
	d, err := NewDriver("", "any_cpf", "any_name", "any_nickname", "234325", s.Address)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), d)
	assert.NotNil(s.T(), d.GetAddress())
	assert.NotNil(s.T(), d.GetID())
	assert.Equal(s.T(), d.GetCPF(), "any_cpf")
	assert.Equal(s.T(), d.GetName(), "any_name")
	assert.Equal(s.T(), d.GetNickName(), "any_nickname")
	assert.Equal(s.T(), d.GetPhone(), "234325")
}

func (s *DriverTestSuite) TestNewDriverWithPassedID() {
	id := identity.NewID()
	d, err := NewDriver(id.String(), "any_cpf", "any_name", "any_nickname", "234325", s.Address)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), d.GetID(), id)

}

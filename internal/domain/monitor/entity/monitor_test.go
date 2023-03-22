package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MonitorestSuite struct {
	suite.Suite
	Address valueobjects.Address
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MonitorestSuite))
}

func (suite *MonitorestSuite) SetupTest() {
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", "123", "")
	suite.Address = *addr
}

func (s *MonitorestSuite) TestNewMonitor() {
	m, err := NewMonitor("any_name", "any_cpf", "any_phonenumber", s.Address)
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), m.ID)
	assert.NotNil(s.T(), m)
	assert.NotNil(s.T(), m.ID)
	assert.Equal(s.T(), m.Name, "any_name")
	assert.Equal(s.T(), m.CPF, "any_cpf")
	assert.Equal(s.T(), m.PhoneNumber, "any_phonenumber")
	assert.EqualValues(s.T(), m.Address, s.Address)
}

func (s *MonitorestSuite) TestNewMonitor_InvalidName() {
	m, err := NewMonitor("", "any_cpf", "any_phonenumber", s.Address)
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid name")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidCPF() {
	m, err := NewMonitor("any_name", "", "any_phonenumber", s.Address)
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid cpf")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidPhonenumber() {
	m, err := NewMonitor("any_name", "any_cpf", "", s.Address)
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid phonenumber")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidAddress() {
	m, err := NewMonitor("any_name", "any_cpf", "any_phone", valueobjects.Address{})
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "address must be a valid address")
}

func (s *MonitorestSuite) TestMonitor_GetAddress() {
	m, _ := NewMonitor("any_name", "any_cpf", "any_phone", s.Address)
	assert.EqualValues(s.T(), m.Address, s.Address)
}

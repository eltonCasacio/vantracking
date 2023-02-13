package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
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
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", 123)
	suite.Address = *addr
}

func (s *MonitorestSuite) TestNewMonitor() {
	m, err := NewMonitor("any_name", "any_cpf", "any_phonenumber", s.Address, []entity.Passenger{})
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), m)
	assert.NotNil(s.T(), m.GetID())
	assert.Equal(s.T(), m.GetName(), "any_name")
	assert.Equal(s.T(), m.GetCPF(), "any_cpf")
	assert.Equal(s.T(), m.GetPhoneNumber(), "any_phonenumber")
	assert.EqualValues(s.T(), m.address, s.Address)
	assert.Equal(s.T(), len(m.GetPassengers()), 0)
}

func (s *MonitorestSuite) TestNewMonitor_InvalidName() {
	m, err := NewMonitor("", "any_cpf", "any_phonenumber", s.Address, []entity.Passenger{})
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid name")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidCPF() {
	m, err := NewMonitor("any_name", "", "any_phonenumber", s.Address, []entity.Passenger{})
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid cpf")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidPhonenumber() {
	m, err := NewMonitor("any_name", "any_cpf", "", s.Address, []entity.Passenger{})
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid phonenumber")
}

func (s *MonitorestSuite) TestNewMonitor_InvalidAddress() {
	m, err := NewMonitor("any_name", "any_cpf", "any_phone", valueobjects.Address{}, []entity.Passenger{})
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "address must be a valid address")
}

func (s *MonitorestSuite) TestMonitor_GetAddress() {
	m, _ := NewMonitor("any_name", "any_cpf", "any_phone", s.Address, []entity.Passenger{})
	assert.EqualValues(s.T(), m.GetAddress(), s.Address)
}

func (s *MonitorestSuite) TestMonitor_AddPassenger() {
	m, _ := NewMonitor("any_name", "any_cpf", "any_phone", s.Address, []entity.Passenger{})
	passenger, _ := entity.NewPassenger("any_name", "any_nickname", "any_routecode", m.GetID())
	m.AddPassenger(*passenger)
	assert.Equal(s.T(), len(m.GetPassengers()), 1)

	passengers := m.GetPassengers()
	assert.Equal(s.T(), passengers[0].GetName(), "any_name")
	assert.Equal(s.T(), passengers[0].GetNickname(), "any_nickname")
	assert.Equal(s.T(), passengers[0].GetRouteCode(), "any_routecode")
	assert.Equal(s.T(), passengers[0].GetMonitorID(), m.GetID())
}

func (s *MonitorestSuite) TestMonitor_AddInvalidPassenger() {
	m, _ := NewMonitor("any_name", "any_cpf", "any_phone", s.Address, []entity.Passenger{})
	err := m.AddPassenger(entity.Passenger{})
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), len(m.GetPassengers()), 0)

}

package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
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
	addr, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "123", "123", "", "", "")
	suite.Address = *addr
}

func (s *DriverTestSuite) TestNewDriver_ErrorAddresEmpty() {
	d, err := NewDriver("234534534", "any_name", "", "24534534", valueobjects.Address{})
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
}

func (s *DriverTestSuite) TestNewDriver_ErrorCPFInvalid() {
	d, err := NewDriver("", "any_name", "", "24534534", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "cpf invalid")
}

func (s *DriverTestSuite) TestNewDriver_ErrorNameInvalid() {
	d, err := NewDriver("any_cpf", "", "", "24534534", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "name invalid")
}

func (s *DriverTestSuite) TestNewDriver_ErrorPhoneInvalid() {
	d, err := NewDriver("any_cpf", "any_name", "", "", s.Address)
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), d)
	assert.EqualError(s.T(), err, "phone invalid")
}

func (s *DriverTestSuite) TestNewDriver() {
	d, err := NewDriver("any_cpf", "any_name", "234325", "", s.Address)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), d)
	assert.NotNil(s.T(), d.Address)
	assert.NotNil(s.T(), d.ID)
	assert.Equal(s.T(), d.CPF, "any_cpf")
	assert.Equal(s.T(), d.Name, "any_name")
	assert.Equal(s.T(), d.Phone, "234325")
}

func (s *DriverTestSuite) TestChangeName() {
	d, _ := NewDriver("any_cpf", "any_name", "234325", "", s.Address)
	err := d.ChangeName("other_name")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), d.Name, "other_name")

	err = d.ChangeName("")
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid name")
}

func (s *DriverTestSuite) TestChangeNickName() {
	d, _ := NewDriver("any_cpf", "any_name", "234325", "", s.Address)
	assert.Equal(s.T(), d.Nickname, "")
	d.ChangeNickname("new_nickname")
	assert.Equal(s.T(), d.Nickname, "new_nickname")
}

func (s *DriverTestSuite) TestChangeCPF() {
	d, _ := NewDriver("3232323232", "any_name", "234325", "", s.Address)
	assert.Equal(s.T(), d.CPF, "3232323232")
	d.ChangeCPF("235234456")
	assert.Equal(s.T(), d.CPF, "235234456")

	err := d.ChangeCPF("")
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid cpf")
}

func (s *DriverTestSuite) TestChangePhone() {
	d, _ := NewDriver("3232323232", "any_name", "234325", "", s.Address)
	assert.Equal(s.T(), d.Phone, "234325")
	d.ChangePhone("2222222")
	assert.Equal(s.T(), d.Phone, "2222222")

	err := d.ChangePhone("")
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid phone")
}

func (s *DriverTestSuite) TestChangeAddress() {
	d, _ := NewDriver("3232323232", "any_name", "234325", "", s.Address)
	assert.Equal(s.T(), d.Address, s.Address)

	newAddress, _ := valueobjects.NewAddress("any_uf", "any_city", "any_street", "777", "13277777", "", "", "")
	d.ChangeAddress(*newAddress)
	assert.NotEqual(s.T(), d.Address, s.Address)
	assert.Equal(s.T(), d.Address, *newAddress)

	err := d.ChangeAddress(valueobjects.Address{})
	assert.NotNil(s.T(), err)
	assert.EqualError(s.T(), err, "invalid address")
}

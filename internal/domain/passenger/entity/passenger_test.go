package passenger

import (
	"testing"

	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
)

func TestNewPassenger(t *testing.T) {
	id := identity.NewID()
	p, err := NewPassenger("any_name", "any_routecode", "nickname", "any_school", id)
	assert.Nil(t, err)
	assert.Equal(t, p.MonitorID, id)
	assert.Equal(t, p.Name, "any_name")
	assert.Equal(t, p.RouteCode, "any_routecode")
}

func TestNewPassenge_InvalidName(t *testing.T) {
	p, err := NewPassenger("", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid name")
}

func TestNewPassenge_InvalidRouteCode(t *testing.T) {
	p, err := NewPassenger("any_name", "", "nickname", "any_school", identity.NewID())
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid route code")
}

func TestSetNickname(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Equal(t, p.Nickname, "")

	p.SetNickname("other_nickname")
	assert.Equal(t, p.Nickname, "other_nickname")
}

func TestSetName(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())

	p.SetName("other_name")
	assert.Equal(t, p.Name, "other_name")
}

func TestSetRouteCode(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Equal(t, p.RouteCode, "any_routecode")

	p.SetRouteCode("other_routecode")
	assert.Equal(t, p.RouteCode, "other_routecode")
}

func TestSetInvalidRouteCode(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Equal(t, p.RouteCode, "any_routecode")

	err := p.SetRouteCode("")
	assert.EqualError(t, err, "invalid route code")
}

func TestSetInvalidName(t *testing.T) {
	p, err := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Nil(t, err)

	err = p.SetName("")
	assert.EqualError(t, err, "invalid name")
}

func TestChangeGoNoGo(t *testing.T) {
	p, err := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Nil(t, err)
	assert.Equal(t, p.Goes, true)
	assert.Equal(t, p.Comesback, true)

	p.ChangeGoNoGo(false, true)
	assert.Equal(t, p.Goes, false)
	assert.Equal(t, p.Comesback, true)

	p.ChangeGoNoGo(false, false)
	assert.Equal(t, p.Goes, false)
	assert.Equal(t, p.Comesback, false)
}

func TestIsPassengerConfirmed(t *testing.T) {
	p, err := NewPassenger("any_name", "any_routecode", "nickname", "any_school", identity.NewID())
	assert.Nil(t, err)
	assert.Equal(t, p.IsRegisterConfirmed(), false)

	p.ConfirmRegister(true)
	assert.Equal(t, p.IsRegisterConfirmed(), true)
}

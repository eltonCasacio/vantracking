package entity

import (
	"testing"

	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
)

func TestNewPassenger(t *testing.T) {
	id := identity.NewID()
	p, err := NewPassenger("any_name", "any_nickname", "any_routecode", id)
	assert.Nil(t, err)
	assert.Equal(t, p.GetMonitorID(), id)
	assert.Equal(t, p.GetName(), "any_name")
	assert.Equal(t, p.GetNickname(), "any_nickname")
	assert.Equal(t, p.GetRouteCode(), "any_routecode")
}

func TestNewPassenge_InvalidName(t *testing.T) {
	p, err := NewPassenger("", "any_nickname", "any_routecode", identity.NewID())
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid name")
}

func TestNewPassenge_InvalidRouteCode(t *testing.T) {
	p, err := NewPassenger("any_name", "any_nickname", "", identity.NewID())
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid route code")
}

func TestSetNickname(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_nickname", "any_routecode", identity.NewID())
	assert.Equal(t, p.nickname, "any_nickname")

	p.SetNickname("other_nickname")
	assert.Equal(t, p.GetNickname(), "other_nickname")
}

func TestSetRouteCode(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_nickname", "any_routecode", identity.NewID())
	assert.Equal(t, p.GetRouteCode(), "any_routecode")

	p.SetRouteCode("other_routecode")
	assert.Equal(t, p.GetRouteCode(), "other_routecode")
}

func TestSetInvalidRouteCode(t *testing.T) {
	p, _ := NewPassenger("any_name", "any_nickname", "any_routecode", identity.NewID())
	assert.Equal(t, p.GetRouteCode(), "any_routecode")

	err := p.SetRouteCode("")
	assert.EqualError(t, err, "invalid route code")
}

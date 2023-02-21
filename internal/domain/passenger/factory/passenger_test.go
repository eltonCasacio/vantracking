package passenger

import (
	"testing"

	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := PassengerFactory()
	psg := PassengerFactory()
	assert.Equal(t, p, psg)

	input := NewPassengerInputDTO{
		Name:      "any_name",
		RouteCode: "any_route_code",
		MonitorID: identity.NewID().String(),
	}
	passenger, err := p.NewPassenger(input)
	assert.Nil(t, err)
	assert.Equal(t, passenger.Name, input.Name)
}

func TestNew_InvalidMonitorID(t *testing.T) {
	p := PassengerFactory()

	input := NewPassengerInputDTO{
		Name:      "any_name",
		RouteCode: "any_route_code",
		MonitorID: "",
	}
	passenger, err := p.NewPassenger(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
}

func TestNew_InvalidParams(t *testing.T) {
	p := PassengerFactory()

	input := NewPassengerInputDTO{
		Name:      "",
		RouteCode: "any_route_code",
		MonitorID: identity.NewID().String(),
	}
	passenger, err := p.NewPassenger(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
	assert.EqualError(t, err, "invalid name")
}

func TestInstance(t *testing.T) {
	p := PassengerFactory()

	id := identity.NewID().String()
	input := PassengerInputDTO{
		ID:                id,
		Name:              "name",
		Nickname:          "nickname",
		RouteCode:         "any_route_code",
		Goes:              true,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         identity.NewID().String(),
	}
	passenger, err := p.Instance(input)
	assert.Nil(t, err)
	assert.NotNil(t, passenger)
	assert.Equal(t, passenger.ID.String(), id)
	assert.True(t, passenger.Goes)
	assert.False(t, passenger.Comesback)
}

func TestInstance_InvalidID(t *testing.T) {
	p := PassengerFactory()

	input := PassengerInputDTO{
		ID:                "",
		Name:              "name",
		Nickname:          "nickname",
		RouteCode:         "any_route_code",
		Goes:              true,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         identity.NewID().String(),
	}
	passenger, err := p.Instance(input)
	assert.Nil(t, passenger)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid id")
}

func TestInstance_InvalidMonitorID(t *testing.T) {
	p := PassengerFactory()

	input := PassengerInputDTO{
		ID:                identity.NewID().String(),
		Name:              "name",
		Nickname:          "nickname",
		RouteCode:         "any_route_code",
		Goes:              true,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         "",
	}
	passenger, err := p.Instance(input)
	assert.Nil(t, passenger)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid monitor id")
}

func TestInstance_InvalidParams(t *testing.T) {
	p := PassengerFactory()

	input := PassengerInputDTO{
		ID:                identity.NewID().String(),
		Name:              "",
		Nickname:          "nickname",
		RouteCode:         "any_route_code",
		Goes:              true,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         identity.NewID().String(),
	}
	passenger, err := p.Instance(input)
	assert.Nil(t, passenger)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid name")
}

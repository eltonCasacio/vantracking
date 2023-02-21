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

	input := PassengerInputDTO{
		Name:      "any_name",
		Nickname:  "any_nickname",
		RouteCode: "any_route_code",
		MonitorID: identity.NewID().String(),
	}
	passenger, err := p.New(input)
	assert.Nil(t, err)
	assert.Equal(t, passenger.Name, input.Name)
}

func TestNew_InvalidMonitorID(t *testing.T) {
	p := PassengerFactory()

	input := PassengerInputDTO{
		Name:      "any_name",
		Nickname:  "any_nickname",
		RouteCode: "any_route_code",
		MonitorID: "",
	}
	passenger, err := p.New(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
}

func TestNew_InvalidParams(t *testing.T) {
	p := PassengerFactory()

	input := PassengerInputDTO{
		Name:      "",
		Nickname:  "any_nickname",
		RouteCode: "any_route_code",
		MonitorID: identity.NewID().String(),
	}
	passenger, err := p.New(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
	assert.EqualError(t, err, "invalid name")
}

func TestCreateInstance(t *testing.T) {
	p := PassengerFactory()

	id := identity.NewID().String()
	monitorid := identity.NewID().String()

	input := PassengerInputDTO{
		ID:                id,
		Name:              "any_name",
		Nickname:          "any_nickname",
		RouteCode:         "any_route_code",
		Goes:              true,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         monitorid,
	}
	passenger, err := p.CreateInstance(input)
	assert.Nil(t, err)
	assert.Equal(t, passenger.ID.String(), id)
	assert.Equal(t, passenger.Goes, input.Goes)
	assert.Equal(t, passenger.Comesback, input.Comesback)
	assert.Equal(t, passenger.RegisterConfirmed, input.RegisterConfirmed)
}

func TestCreateInstance_InvalidID(t *testing.T) {
	p := PassengerFactory()

	id := ""
	monitorid := identity.NewID().String()

	input := PassengerInputDTO{
		ID:                id,
		Name:              "any_name",
		Nickname:          "any_nickname",
		RouteCode:         "any_route_code",
		Goes:              false,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         monitorid,
	}
	passenger, err := p.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
	assert.Equal(t, input.ID, "")
}

func TestCreateInstance_InvalidMonitorID(t *testing.T) {
	p := PassengerFactory()

	id := identity.NewID().String()
	monitorid := ""

	input := PassengerInputDTO{
		ID:                id,
		Name:              "any_name",
		Nickname:          "any_nickname",
		RouteCode:         "any_route_code",
		Goes:              false,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         monitorid,
	}
	passenger, err := p.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
	assert.Equal(t, input.ID, id)
	assert.Equal(t, input.MonitorID, "")
}

func TestCreateInstance_InvalidParams(t *testing.T) {
	p := PassengerFactory()

	id := identity.NewID().String()
	monitorid := identity.NewID().String()

	input := PassengerInputDTO{
		ID:                id,
		Name:              "",
		Nickname:          "any_nickname",
		RouteCode:         "any_route_code",
		Goes:              false,
		Comesback:         false,
		RegisterConfirmed: true,
		MonitorID:         monitorid,
	}
	passenger, err := p.CreateInstance(input)
	assert.NotNil(t, err)
	assert.Nil(t, passenger)
	assert.Equal(t, input.ID, id)
	assert.Equal(t, input.MonitorID, monitorid)
	assert.Equal(t, input.Name, "")
}

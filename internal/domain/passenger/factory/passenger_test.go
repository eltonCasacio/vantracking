package factory

import (
	"testing"

	"github.com/eltoncasacio/vantracking/pkg/identity"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	p := PassengerFactory()
	psg := PassengerFactory()
	assert.Equal(t, p, psg)

	input := PassengerInputDTO{
		Name:      "any_name",
		Nickname:  "any_nickname",
		RouteCode: "any_route_code",
		MonitorID: identity.NewID(),
	}
	passenger, err := p.Create(input)
	assert.Nil(t, err)
	assert.Equal(t, passenger.GetName(), input.Name)
}

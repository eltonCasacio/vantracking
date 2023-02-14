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
		name:      "any_name",
		nickname:  "any_nickname",
		routeCode: "any_route_code",
		monitorID: identity.NewID(),
	}
	passenger, err := p.Create(input)
	assert.Nil(t, err)
	assert.Equal(t, passenger.GetName(), input.name)
}

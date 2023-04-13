package valueobjects

import (
	"errors"

	"github.com/eltoncasacio/vantracking/internal/domain/shared/repository"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type DeviceRepositoryInterface interface {
	repository.RepositoryInterface[Device]
}

type Device struct {
	Token     string
	MonitorID identity.ID
}

func NewDevice(token string, monitorID identity.ID) (*Device, error) {
	addr := &Device{
		Token:     token,
		MonitorID: monitorID,
	}

	err := addr.IsValid()
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *Device) IsValid() error {
	if a.Token == "" {
		return errors.New("token is mandatory")
	}
	if a.MonitorID.String() == "" {
		return errors.New("monitor id is mandatory")
	}
	return nil
}

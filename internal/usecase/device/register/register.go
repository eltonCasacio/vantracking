package device

import (
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type RegisterDeviceUseCase struct {
	repository vo.DeviceRepositoryInterface
}

func NewUseCase(repository vo.DeviceRepositoryInterface) *RegisterDeviceUseCase {
	return &RegisterDeviceUseCase{
		repository: repository,
	}
}

func (u *RegisterDeviceUseCase) Register(input *DeviceInputDTO) error {
	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return err
	}

	deviceInput, err := vo.NewDevice(input.Token, monitorID)
	if err != nil {
		return err
	}

	found, err := u.repository.FindByMonitorID(input.MonitorID)
	if err != nil {
		return err
	}

	if found.Token == input.Token {
		return nil
	}

	if found.Token != input.Token {
		err = u.repository.Update(deviceInput)
		if err != nil {
			return err
		}
		return nil
	}

	err = u.repository.Create(deviceInput)
	if err != nil {
		return err
	}
	return nil
}

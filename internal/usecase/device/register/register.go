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

func (u *RegisterDeviceUseCase) Register(input *DriverInputDTO) error {

	monitorID, err := identity.ParseID(input.MonitorID)
	if err != nil {
		return err
	}

	driverInput, err := vo.NewDevice(input.Token, monitorID)
	if err != nil {
		return err
	}

	err = u.repository.Create(driverInput)
	if err != nil {
		return err
	}
	return nil
}

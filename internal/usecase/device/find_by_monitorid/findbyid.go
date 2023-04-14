package device

import (
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type findByMonitorIDUseCase struct {
	repository vo.DeviceRepositoryInterface
}

func NewUseCase(repository vo.DeviceRepositoryInterface) *findByMonitorIDUseCase {
	return &findByMonitorIDUseCase{
		repository: repository,
	}
}

func (cd *findByMonitorIDUseCase) Find(monitorID string) (DeviceOutput, error) {
	device, err := cd.repository.FindByMonitorID(monitorID)
	if err != nil {
		return DeviceOutput{}, err
	}

	output := DeviceOutput{
		Token:     device.Token,
		MonitorID: device.MonitorID.String(),
	}
	return output, nil
}

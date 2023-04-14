package device

type DeviceInput struct {
	MonitorID string `json:"monitorID"`
}

type DeviceOutput struct {
	Token     string `json:"token"`
	MonitorID string `json:"monitorID"`
}

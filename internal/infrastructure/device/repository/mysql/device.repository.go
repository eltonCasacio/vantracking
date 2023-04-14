package device

import (
	"database/sql"
	"errors"

	entity "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
	"github.com/eltoncasacio/vantracking/pkg/identity"
)

type DeviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (d *DeviceRepository) Create(device *entity.Device) error {
	if err := device.IsValid(); err != nil {
		return err
	}

	query := `INSERT INTO devices(token, monitor_id) values(?,?)`
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		device.Token,
		device.MonitorID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceRepository) FindByMonitorID(monitorID string) (*entity.Device, error) {
	if monitorID == "" {
		return nil, errors.New("monitorID is required")
	}

	stmt, err := d.db.Prepare("SELECT token, monitor_id FROM devices WHERE monitor_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var token, idMonitor string
	err = stmt.QueryRow(monitorID).Scan(&token, &idMonitor)
	if err != nil {
		return nil, err
	}

	parsedMonitorID, err := identity.ParseID(idMonitor)
	if err != nil {
		return nil, err
	}
	output := entity.Device{Token: token, MonitorID: parsedMonitorID}
	return &output, nil
}

func (d *DeviceRepository) Update(device *entity.Device) error {
	stmt, err := d.db.Prepare("update devices set token=? WHERE monitor_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		device.Token,
		device.MonitorID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceRepository) FindByID(id string) (*entity.Device, error) {
	return nil, nil
}

func (d *DeviceRepository) Delete(id string) error {
	return nil
}

func (d *DeviceRepository) FindAll() ([]entity.Device, error) {
	return nil, nil
}

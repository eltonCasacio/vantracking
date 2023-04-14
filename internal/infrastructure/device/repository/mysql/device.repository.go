package device

import (
	"database/sql"

	entity "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
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

func (d *DeviceRepository) FindByID(id string) (*entity.Device, error) {
	// if id == "" {
	// 	return nil, errors.New("id is required")
	// }

	// stmt, err := d.db.Prepare("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement, latitude, longitude FROM drivers WHERE id = ? and active = true")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// var driverInput factory.CreateInstanceDriverInputDTO
	// err = stmt.QueryRow(id).Scan(
	// 	&driverInput.ID,
	// 	&driverInput.CPF,
	// 	&driverInput.Name,
	// 	&driverInput.Nickname,
	// 	&driverInput.Phone,
	// 	&driverInput.UF,
	// 	&driverInput.City,
	// 	&driverInput.Street,
	// 	&driverInput.Number,
	// 	&driverInput.CEP,
	// 	&driverInput.Complement,
	// 	&driverInput.Latitude,
	// 	&driverInput.Longitude,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// driver, err := factory.DriverFactory().CreateInstance(driverInput)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (d *DeviceRepository) Update(device *entity.Device) error {
	// if err := driver.IsValid(); err != nil {
	// 	return errors.New("invalid driver")
	// }
	// query := "UPDATE drivers SET cpf = ?, name = ?, nickname = ?, phone = ?, uf = ?, city = ?, street = ?, number = ?, cep = ?, complement = ? , latitude=?, longitude=? WHERE id = ?"
	// stmt, err := d.db.Prepare(query)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()

	// addr := driver.Address

	// _, err = stmt.Exec(
	// 	driver.CPF,
	// 	driver.Name,
	// 	driver.Nickname,
	// 	driver.Phone,
	// 	addr.UF,
	// 	addr.City,
	// 	addr.Street,
	// 	addr.Number,
	// 	addr.CEP,
	// 	addr.Complement,
	// 	addr.Latitude,
	// 	addr.Longitude,
	// 	driver.ID,
	// )
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (d *DeviceRepository) Delete(id string) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("DELETE FROM drivers WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM routes WHERE driver_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (d *DeviceRepository) FindAll() ([]entity.Device, error) {
	// rows, err := d.db.Query("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement, latitude, longitude FROM drivers WHERE active = true")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var drivers []entity.Device
	// for rows.Next() {
	// 	var driverInput factory.CreateInstanceDriverInputDTO
	// 	err := rows.Scan(
	// 		&driverInput.ID,
	// 		&driverInput.CPF,
	// 		&driverInput.Name,
	// 		&driverInput.Nickname,
	// 		&driverInput.Phone,
	// 		&driverInput.UF,
	// 		&driverInput.City,
	// 		&driverInput.Street,
	// 		&driverInput.Number,
	// 		&driverInput.CEP,
	// 		&driverInput.Complement,
	// 		&driverInput.Latitude,
	// 		&driverInput.Longitude,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	d, err := factory.DriverFactory().CreateInstance(driverInput)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	drivers = append(drivers, *d)
	// }
	return nil, nil
}

func (d *DeviceRepository) FindByMonitorID(monitorID string) (*entity.Device, error) {
	// if id == "" {
	// 	return nil, errors.New("id is required")
	// }

	// stmt, err := d.db.Prepare("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement, latitude, longitude FROM drivers WHERE id = ? and active = true")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// var driverInput factory.CreateInstanceDriverInputDTO
	// err = stmt.QueryRow(id).Scan(
	// 	&driverInput.ID,
	// 	&driverInput.CPF,
	// 	&driverInput.Name,
	// 	&driverInput.Nickname,
	// 	&driverInput.Phone,
	// 	&driverInput.UF,
	// 	&driverInput.City,
	// 	&driverInput.Street,
	// 	&driverInput.Number,
	// 	&driverInput.CEP,
	// 	&driverInput.Complement,
	// 	&driverInput.Latitude,
	// 	&driverInput.Longitude,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// driver, err := factory.DriverFactory().CreateInstance(driverInput)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

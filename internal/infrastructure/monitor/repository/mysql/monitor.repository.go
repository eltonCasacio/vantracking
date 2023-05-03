package repository

import (
	"database/sql"
	"errors"

	driver "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	driverFactory "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	"github.com/eltoncasacio/vantracking/internal/domain/monitor/entity"
	"github.com/eltoncasacio/vantracking/internal/domain/monitor/factory"
)

type MonitorRepository struct {
	db *sql.DB
}

func NewMonitorRepository(db *sql.DB) *MonitorRepository {
	return &MonitorRepository{db: db}
}

func (m *MonitorRepository) Create(monitor *entity.Monitor) error {
	stmt, err := m.db.Prepare(`
	INSERT 
	INTO monitors (
		id,
		cpf,
		name,
		phone_number,
		uf,
		city,
		street,
		number,
		cep,
		complement,
		latitude,
		longitude,
		active
	)
	values(?,?,?,?,?,?,?,?,?,?,?,?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	address := monitor.Address

	_, err = stmt.Exec(
		monitor.ID.String(),
		monitor.CPF,
		monitor.Name,
		monitor.PhoneNumber,
		address.UF,
		address.City,
		address.Street,
		address.Number,
		address.CEP,
		address.Complement,
		address.Latitude,
		address.Longitude,
		true,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Update(monitor *entity.Monitor) error {
	stmt, err := m.db.Prepare("update monitors set cpf = ?, name = ?, phone_number = ?, uf = ?, city = ?, street = ?, number = ?, cep  = ?, complement = ? ,latitude=?, longitude=? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	address := monitor.Address

	_, err = stmt.Exec(
		monitor.CPF,
		monitor.Name,
		monitor.PhoneNumber,
		address.UF,
		address.City,
		address.Street,
		address.Number,
		address.CEP,
		address.Complement,
		address.Latitude,
		address.Longitude,
		monitor.ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Delete(id string) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := m.db.Prepare("delete from monitors where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM passengers WHERE monitor_id = ?")
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

func (m *MonitorRepository) FindAll() ([]entity.Monitor, error) {
	rows, err := m.db.Query("SELECT id, cpf, name, phone_number, uf, city, street, number, cep, complement,latitude, longitude FROM monitors WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monitors []entity.Monitor
	for rows.Next() {
		inputMonitor := factory.InstanceMonitorInputDTO{}
		err := rows.Scan(
			&inputMonitor.ID,
			&inputMonitor.CPF,
			&inputMonitor.Name,
			&inputMonitor.PhoneNumber,
			&inputMonitor.UF,
			&inputMonitor.City,
			&inputMonitor.Street,
			&inputMonitor.Number,
			&inputMonitor.CEP,
			&inputMonitor.Complement,
			&inputMonitor.Latitude,
			&inputMonitor.Longitude,
		)
		if err != nil {
			return nil, err
		}

		newMonitor, err := factory.MonitorFactory().Instance(inputMonitor)
		if err != nil {
			return nil, err
		}
		monitors = append(monitors, *newMonitor)
	}
	return monitors, nil
}

func (m *MonitorRepository) FindByID(id string) (*entity.Monitor, error) {
	stmt, err := m.db.Prepare("SELECT id, cpf, name, phone_number, uf, city, street, number, cep, complement,latitude, longitude FROM monitors WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inputMonitor := factory.InstanceMonitorInputDTO{}
	rows := stmt.QueryRow(id)
	rows.Scan(
		&inputMonitor.ID,
		&inputMonitor.CPF,
		&inputMonitor.Name,
		&inputMonitor.PhoneNumber,
		&inputMonitor.UF,
		&inputMonitor.City,
		&inputMonitor.Street,
		&inputMonitor.Number,
		&inputMonitor.CEP,
		&inputMonitor.Complement,
		&inputMonitor.Latitude,
		&inputMonitor.Longitude,
	)
	if err != nil {
		return nil, err
	}

	newMonitor, err := factory.MonitorFactory().Instance(inputMonitor)
	if err != nil {
		return nil, err
	}
	return newMonitor, nil
}

func (d *MonitorRepository) FindByCPF(cpf string) (*entity.Monitor, error) {
	stmt, err := d.db.Prepare("SELECT id, cpf, name, phone_number, uf, city, street, number, cep, complement,latitude, longitude FROM monitors WHERE cpf = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inputMonitor := factory.InstanceMonitorInputDTO{}
	err = stmt.QueryRow(cpf).Scan(
		&inputMonitor.ID,
		&inputMonitor.CPF,
		&inputMonitor.Name,
		&inputMonitor.PhoneNumber,
		&inputMonitor.UF,
		&inputMonitor.City,
		&inputMonitor.Street,
		&inputMonitor.Number,
		&inputMonitor.CEP,
		&inputMonitor.Complement,
		&inputMonitor.Latitude,
		&inputMonitor.Longitude,
	)
	if err != nil {
		return nil, err
	}

	driver, err := factory.MonitorFactory().Instance(inputMonitor)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *MonitorRepository) GetDriverByRouteCode(routeCode string) (*driver.Driver, error) {
	if routeCode == "" {
		return nil, errors.New("driver id is required")
	}

	stmt, err := d.db.Prepare(`
	SELECT d.id, d.cpf, d.name, d.nickname, d.phone, d.uf, d.city, d.street, d.number, d.cep, d.complement,latitude, longitude
	FROM drivers as d 
	INNER JOIN routes
	ON code = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var driverInput driverFactory.CreateInstanceDriverInputDTO
	err = stmt.QueryRow(routeCode).Scan(
		&driverInput.ID,
		&driverInput.CPF,
		&driverInput.Name,
		&driverInput.Nickname,
		&driverInput.Phone,
		&driverInput.UF,
		&driverInput.City,
		&driverInput.Street,
		&driverInput.Number,
		&driverInput.CEP,
		&driverInput.Complement,
		&driverInput.Latitude,
		&driverInput.Longitude,
	)
	if err != nil {
		return nil, err
	}

	driver, err := driverFactory.DriverFactory().CreateInstance(driverInput)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

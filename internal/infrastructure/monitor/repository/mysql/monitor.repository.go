package repository

import (
	"database/sql"

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
	stmt, err := m.db.Prepare("insert into monitors(id, cpf, name, phone_number, uf, city, street, number, cep , active) values(?,?,?,?,?,?,?,?,?,?)")
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
		true,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Update(monitor *entity.Monitor) error {
	stmt, err := m.db.Prepare("update monitors set cpf = ?, name = ?, phone_number = ?, uf = ?, city = ?, street = ?, number = ?, cep  = ? WHERE id = ?")
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
		monitor.ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) Delete(id string) error {
	stmt, err := m.db.Prepare("delete from monitors where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MonitorRepository) FindAll() ([]entity.Monitor, error) {
	rows, err := m.db.Query("SELECT id, cpf, name, phone_number, uf, city, street, number, cep FROM monitors WHERE active = true")
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
	stmt, err := m.db.Prepare("SELECT id, cpf, name, phone_number, uf, city, street, number, cep FROM monitors WHERE id = ? and active = true")
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
	stmt, err := d.db.Prepare("SELECT id, cpf, name, phone_number, uf, city, street, number, cep FROM monitors WHERE cpf = ? and active = true")
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

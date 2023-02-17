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

	address := monitor.GetAddress()

	model := MonitorModel{
		id:          monitor.GetID().String(),
		name:        monitor.GetName(),
		cpf:         monitor.GetCPF(),
		phoneNumber: monitor.GetPhoneNumber(),
		uf:          address.GetUF(),
		city:        address.GetCity(),
		street:      address.GetStreet(),
		number:      address.GetNumber(),
		cep:         address.GetCEP(),
		active:      true,
	}
	_, err = stmt.Exec(
		model.id,
		model.cpf,
		model.name,
		model.phoneNumber,
		model.uf,
		model.city,
		model.street,
		model.number,
		model.cep,
		model.active,
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

	address := monitor.GetAddress()

	_, err = stmt.Exec(
		monitor.GetCPF(),
		monitor.GetName(),
		monitor.GetPhoneNumber(),
		address.GetUF(),
		address.GetCity(),
		address.GetStreet(),
		address.GetNumber(),
		address.GetCEP(),
		monitor.GetID().String(),
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
	rows, err := m.db.Query("SELECT * FROM monitors WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monitors []entity.Monitor
	for rows.Next() {
		var model MonitorModel
		err := rows.Scan(
			&model.id,
			&model.cpf,
			&model.name,
			&model.phoneNumber,
			&model.uf,
			&model.city,
			&model.street,
			&model.number,
			&model.cep,
			&model.active,
		)
		if err != nil {
			return nil, err
		}

		inputMonitor := factory.CreateMonitorInputDTO{
			ID:          model.id,
			Name:        model.name,
			CPF:         model.cpf,
			PhoneNumber: model.phoneNumber,
			UF:          model.uf,
			City:        model.city,
			Street:      model.street,
			Number:      model.number,
			CEP:         model.cep,
		}

		newMonitor, err := factory.MonitorFactory().Create(inputMonitor)
		if err != nil {
			return nil, err
		}
		monitors = append(monitors, *newMonitor)
	}
	return monitors, nil
}

func (m *MonitorRepository) FindByID(id string) (*entity.Monitor, error) {
	stmt, err := m.db.Prepare("SELECT * FROM monitors WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var model MonitorModel
	rows := stmt.QueryRow(id)
	rows.Scan(
		&model.id,
		&model.cpf,
		&model.name,
		&model.phoneNumber,
		&model.uf,
		&model.city,
		&model.street,
		&model.number,
		&model.cep,
		&model.active,
	)
	if err != nil {
		return nil, err
	}

	inputMonitor := factory.CreateMonitorInputDTO{
		ID:          model.id,
		Name:        model.name,
		CPF:         model.cpf,
		PhoneNumber: model.phoneNumber,
		UF:          model.uf,
		City:        model.city,
		Street:      model.street,
		Number:      model.number,
		CEP:         model.cep,
	}

	newMonitor, err := factory.MonitorFactory().Create(inputMonitor)
	if err != nil {
		return nil, err
	}
	return newMonitor, nil
}

func (d *MonitorRepository) FindByCPF(cpf string) (*entity.Monitor, error) {
	stmt, err := d.db.Prepare("SELECT * FROM monitors WHERE cpf = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var model MonitorModel
	err = stmt.QueryRow(cpf).Scan(
		&model.id,
		&model.cpf,
		&model.name,
		&model.phoneNumber,
		&model.uf,
		&model.city,
		&model.street,
		&model.number,
		&model.cep,
		&model.active,
	)
	if err != nil {
		return nil, err
	}

	inputMonitor := factory.CreateMonitorInputDTO{
		ID:          model.id,
		Name:        model.name,
		CPF:         model.cpf,
		PhoneNumber: model.phoneNumber,
		UF:          model.uf,
		City:        model.city,
		Street:      model.street,
		Number:      model.number,
		CEP:         model.cep,
	}
	driver, err := factory.MonitorFactory().Create(inputMonitor)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

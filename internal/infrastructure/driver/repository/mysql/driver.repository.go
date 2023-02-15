package repository

import (
	"database/sql"
	"errors"

	entity "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	factory "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
)

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (d *DriverRepository) Create(driver *entity.Driver) error {
	if err := driver.IsValid(); err != nil {
		return err
	}
	query := `INSERT INTO drivers(id, cpf, name, nickname, phone, uf, city, street, number, cep) values(?,?,?,?,?,?,?,?,?,?)`
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	addr := driver.GetAddress()

	_, err = stmt.Exec(
		driver.GetID().String(),
		driver.GetCPF(),
		driver.GetName(),
		driver.GetNickName(),
		driver.GetPhone(),
		addr.GetUF(),
		addr.GetCity(),
		addr.GetStreet(),
		addr.GetNumber(),
		addr.GetCEP(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) FindByID(id string) (*entity.Driver, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	stmt, err := d.db.Prepare("SELECT * FROM drivers WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var model DriverModel
	err = stmt.QueryRow(id).Scan(
		&model.id,
		&model.cpf,
		&model.name,
		&model.nickname,
		&model.phone,
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

	driverInput := factory.DriverInputDTO{
		ID:       model.id,
		CPF:      model.cpf,
		Name:     model.name,
		Nickname: model.nickname,
		Phone:    model.phone,
		UF:       model.uf,
		City:     model.city,
		Street:   model.street,
		Number:   model.number,
		CEP:      model.cep,
	}
	driver, err := factory.DriverFactory().Create(driverInput)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *DriverRepository) Update(driver *entity.Driver) error {
	if err := driver.IsValid(); err != nil {
		return errors.New("invalid driver")
	}
	query := "UPDATE drivers SET cpf = ?, name = ?, nickname = ?, phone = ?, uf = ?, city = ?, street = ?, number = ?, cep = ? WHERE id = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	addr := driver.GetAddress()

	_, err = stmt.Exec(
		driver.GetCPF(),
		driver.GetName(),
		driver.GetNickName(),
		driver.GetPhone(),
		addr.GetUF(),
		addr.GetCity(),
		addr.GetStreet(),
		addr.GetNumber(),
		addr.GetCEP(),
		driver.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) Delete(id string) error {
	stmt, err := d.db.Prepare("UPDATE drivers SET active = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(false, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) FindAll() ([]entity.Driver, error) {
	rows, err := d.db.Query("SELECT * FROM drivers WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []entity.Driver
	for rows.Next() {
		var model DriverModel
		err := rows.Scan(
			&model.id,
			&model.cpf,
			&model.name,
			&model.nickname,
			&model.phone,
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

		driverInput := factory.DriverInputDTO{
			ID:       model.id,
			CPF:      model.cpf,
			Name:     model.name,
			Nickname: model.nickname,
			Phone:    model.phone,
			UF:       model.uf,
			City:     model.city,
			Street:   model.street,
			Number:   model.number,
			CEP:      model.cep,
		}

		d, err := factory.DriverFactory().Create(driverInput)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, *d)
	}
	return drivers, nil
}

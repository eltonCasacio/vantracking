package repository

import (
	"database/sql"
	"errors"

	entity "github.com/eltoncasacio/vantracking/internal/domain/driver/entity"
	factory "github.com/eltoncasacio/vantracking/internal/domain/driver/factory"
	route "github.com/eltoncasacio/vantracking/internal/domain/route"
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
	query := `INSERT INTO drivers(id, cpf, name, nickname, phone, uf, city, street, number, cep, complement) values(?, ?,?,?,?,?,?,?,?,?,?)`
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	addr := driver.Address

	_, err = stmt.Exec(
		driver.ID.String(),
		driver.CPF,
		driver.Name,
		driver.Nickname,
		driver.Phone,
		addr.UF,
		addr.City,
		addr.Street,
		addr.Number,
		addr.CEP,
		addr.Complement,
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

	stmt, err := d.db.Prepare("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement FROM drivers WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var driverInput factory.CreateInstanceDriverInputDTO
	err = stmt.QueryRow(id).Scan(
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
	)
	if err != nil {
		return nil, err
	}

	driver, err := factory.DriverFactory().CreateInstance(driverInput)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *DriverRepository) Update(driver *entity.Driver) error {
	if err := driver.IsValid(); err != nil {
		return errors.New("invalid driver")
	}
	query := "UPDATE drivers SET cpf = ?, name = ?, nickname = ?, phone = ?, uf = ?, city = ?, street = ?, number = ?, cep = ?, complement = ? WHERE id = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	addr := driver.Address

	_, err = stmt.Exec(
		driver.CPF,
		driver.Name,
		driver.Nickname,
		driver.Phone,
		addr.UF,
		addr.City,
		addr.Street,
		addr.Number,
		addr.CEP,
		addr.Complement,
		driver.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) Delete(id string) error {
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

func (d *DriverRepository) FindAll() ([]entity.Driver, error) {
	rows, err := d.db.Query("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement FROM drivers WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []entity.Driver
	for rows.Next() {
		var driverInput factory.CreateInstanceDriverInputDTO
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, err
		}

		d, err := factory.DriverFactory().CreateInstance(driverInput)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, *d)
	}
	return drivers, nil
}

func (d *DriverRepository) FindByCPF(cpf string) (*entity.Driver, error) {
	stmt, err := d.db.Prepare("SELECT id, cpf, name, nickname, phone, uf, city, street, number, cep, complement FROM drivers WHERE cpf = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var driverInput factory.CreateInstanceDriverInputDTO
	err = stmt.QueryRow(cpf).Scan(
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
	)
	if err != nil {
		return nil, err
	}

	driver, err := factory.DriverFactory().CreateInstance(driverInput)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *DriverRepository) CreateRoute(route *route.Route) error {
	if err := route.IsValid(); err != nil {
		return err
	}

	stmt, _ := d.db.Prepare("SELECT id FROM drivers WHERE id = ? and active = true")
	defer stmt.Close()
	var driverID string
	stmt.QueryRow(route.DriverID).Scan(&driverID)
	if driverID != route.DriverID.String() {
		return errors.New("invalid driver ID")
	}

	stmt, _ = d.db.Prepare("SELECT name FROM routes WHERE name = ?")
	defer stmt.Close()
	var name string
	stmt.QueryRow(route.Name).Scan(&name)
	if name == route.Name {
		return errors.New("already exists the same name, choice another name")
	}

	stmt, _ = d.db.Prepare(`INSERT INTO routes(code, name, driver_id, started) values(?,?,?,?)`)
	defer stmt.Close()

	_, err := stmt.Exec(
		route.Code,
		route.Name,
		route.DriverID,
		route.Started,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) DeleteRoute(code string) error {
	stmt, err := d.db.Prepare("DELETE FROM routes WHERE code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(code)
	if err != nil {
		return err
	}
	return nil
}

func (d *DriverRepository) Routes(driverID string) ([]route.Route, error) {
	rows, err := d.db.Query("SELECT name, code, driver_id, started FROM routes WHERE driver_id = ?", driverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []route.Route
	for rows.Next() {
		var input route.Route
		err := rows.Scan(
			&input.Name,
			&input.Code,
			&input.DriverID,
			&input.Started,
		)
		if err != nil {
			return nil, err
		}

		routes = append(routes, input)
	}
	return routes, nil
}

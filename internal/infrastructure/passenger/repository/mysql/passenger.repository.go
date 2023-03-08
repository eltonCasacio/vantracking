package repository

import (
	"database/sql"
	"errors"

	e "github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	f "github.com/eltoncasacio/vantracking/internal/domain/passenger/factory"
)

type passengerRepository struct {
	db *sql.DB
}

func NewPassengerRepository(db *sql.DB) *passengerRepository {
	return &passengerRepository{db: db}
}

func (r *passengerRepository) Create(passenger *e.Passenger) error {
	// PARA VERIFICAR SE O MONITOR É VALIDO
	stmt, err := r.db.Prepare("SELECT id FROM monitors WHERE id = ? and active = true")
	if err != nil {
		return err
	}
	defer stmt.Close()
	var id string
	stmt.QueryRow(passenger.MonitorID.String()).Scan(&id)
	if id == "" {
		return errors.New("monitor id is invalid")
	}

	// VERIFICAR SE JA EXISTE
	stmt, err = r.db.Prepare("SELECT name FROM passengers WHERE name = ? and monitor_id = ? and active = true")
	if err != nil {
		return err
	}
	defer stmt.Close()
	var name string
	stmt.QueryRow(passenger.Name, passenger.MonitorID.String()).Scan(&name)
	if name != "" {
		return errors.New("passenger already exists")
	}

	//CRIA PASSAGEIRO
	stmt, err = r.db.Prepare("INSERT INTO passengers (id, name, nickname, route_code, monitor_id) values(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		passenger.ID.String(),
		passenger.Name,
		passenger.Nickname,
		passenger.RouteCode,
		passenger.MonitorID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *passengerRepository) Update(passenger *e.Passenger) error {
	stmt, err := r.db.Prepare("UPDATE passengers SET name = ?, nickname = ?, route_code = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		passenger.Name,
		passenger.Nickname,
		passenger.RouteCode,
		passenger.ID.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *passengerRepository) FindAll() ([]e.Passenger, error) {
	rows, err := r.db.Query("SELECT id, name, nickname, route_code, goes, comesback, register_confirmed, monitor_id FROM passengers WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Passengers []e.Passenger
	for rows.Next() {
		inputPassenger := f.PassengerInputDTO{}
		err := rows.Scan(
			&inputPassenger.ID,
			&inputPassenger.Name,
			&inputPassenger.Nickname,
			&inputPassenger.RouteCode,
			&inputPassenger.Goes,
			&inputPassenger.Comesback,
			&inputPassenger.RegisterConfirmed,
			&inputPassenger.MonitorID,
		)
		if err != nil {
			return nil, err
		}

		newPassenger, err := f.PassengerFactory().Instance(inputPassenger)

		if err != nil {
			return nil, err
		}
		Passengers = append(Passengers, *newPassenger)
	}
	return Passengers, nil
}

func (r *passengerRepository) FindByID(id string) (*e.Passenger, error) {
	stmt, err := r.db.Prepare("SELECT id, name, nickname, route_code, goes, comesback, register_confirmed, monitor_id FROM passengers WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	inputPassenger := f.PassengerInputDTO{}
	rows := stmt.QueryRow(id)
	rows.Scan(
		&inputPassenger.ID,
		&inputPassenger.Name,
		&inputPassenger.Nickname,
		&inputPassenger.RouteCode,
		&inputPassenger.Goes,
		&inputPassenger.Comesback,
		&inputPassenger.RegisterConfirmed,
		&inputPassenger.MonitorID,
	)
	if err != nil {
		return nil, err
	}

	newPassenger, err := f.PassengerFactory().Instance(inputPassenger)
	if err != nil {
		return nil, err
	}
	return newPassenger, nil
}

func (r *passengerRepository) ListNotConfirmedPassengers() ([]e.Passenger, error) {
	rows, err := r.db.Query("SELECT id, name, nickname, route_code, goes, comesback, register_confirmed, monitor_id FROM passengers WHERE register_confirmed = false  active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []e.Passenger
	for rows.Next() {
		inputPassenger := f.PassengerInputDTO{}
		err := rows.Scan(
			&inputPassenger.ID,
			&inputPassenger.Name,
			&inputPassenger.Nickname,
			&inputPassenger.RouteCode,
			&inputPassenger.Goes,
			&inputPassenger.Comesback,
			&inputPassenger.MonitorID,
		)
		if err != nil {
			return nil, err
		}

		newPassenger, err := f.PassengerFactory().Instance(inputPassenger)
		if err != nil {
			return nil, err
		}
		passengers = append(passengers, *newPassenger)
	}

	return passengers, nil
}

func (r *passengerRepository) Delete(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM passengers WHERE id = ?")
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

func (r *passengerRepository) ConfirmPassengerRegister(id string, confirm bool) error {
	stmt, err := r.db.Prepare("UPDATE passengers SET  register_confirmed = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		confirm,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *passengerRepository) ListByMonitorID(monitorID string) ([]e.Passenger, error) {
	rows, err := r.db.Query("SELECT id, name, nickname, route_code, goes, comesback, monitor_id FROM passengers WHERE monitor_id = ? and  active = true", monitorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []e.Passenger
	for rows.Next() {
		inputPassenger := f.PassengerInputDTO{}
		err := rows.Scan(
			&inputPassenger.ID,
			&inputPassenger.Name,
			&inputPassenger.Nickname,
			&inputPassenger.RouteCode,
			&inputPassenger.Goes,
			&inputPassenger.Comesback,
			&inputPassenger.MonitorID,
		)
		if err != nil {
			return nil, err
		}

		newPassenger, err := f.PassengerFactory().Instance(inputPassenger)
		if err != nil {
			return nil, err
		}
		passengers = append(passengers, *newPassenger)
	}

	return passengers, nil
}

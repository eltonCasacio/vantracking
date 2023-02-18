package repository

import (
	"database/sql"

	"github.com/eltoncasacio/vantracking/internal/domain/passenger/entity"
	"github.com/eltoncasacio/vantracking/internal/domain/passenger/factory"
)

type passengerRepository struct {
	db *sql.DB
}

func NewPassengerRepository(db *sql.DB) *passengerRepository {
	return &passengerRepository{db: db}
}

func (r *passengerRepository) Create(passenger *entity.Passenger) error {
	stmt, err := r.db.Prepare("INSERT INTO passengers (id, name, nickname, route_code, monitor_id) values(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	model := PassengerModel{
		ID:        passenger.GetID().String(),
		Name:      passenger.GetName(),
		Nickname:  passenger.GetNickname(),
		RouteCode: passenger.GetRouteCode(),
		MonitorID: passenger.GetMonitorID().String(),
	}

	_, err = stmt.Exec(
		model.ID,
		model.Name,
		model.Nickname,
		model.RouteCode,
		model.MonitorID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *passengerRepository) Update(passenger *entity.Passenger) error {
	stmt, err := r.db.Prepare("UPDATE passengers SET name = ?, nickname = ?, route_code = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		passenger.GetName(),
		passenger.GetNickname(),
		passenger.GetRouteCode(),
		passenger.GetID().String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *passengerRepository) FindAll() ([]entity.Passenger, error) {
	rows, err := r.db.Query("SELECT * FROM passengers WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Passengers []entity.Passenger
	for rows.Next() {
		var model PassengerModel
		err := rows.Scan(
			&model.ID,
			&model.Name,
			&model.Nickname,
			&model.RouteCode,
			&model.Goes,
			&model.Comesback,
			&model.RegisterConfirmed,
			&model.MonitorID,
			&model.active,
		)
		if err != nil {
			return nil, err
		}

		inputPassenger := factory.PassengerInputDTO{
			ID:                model.ID,
			Name:              model.Name,
			Nickname:          model.Nickname,
			RouteCode:         model.RouteCode,
			Goes:              model.Goes,
			Comesback:         model.Comesback,
			RegisterConfirmed: model.RegisterConfirmed,
			MonitorID:         model.MonitorID,
		}

		newPassenger, err := factory.PassengerFactory().Create(inputPassenger)

		if err != nil {
			return nil, err
		}
		Passengers = append(Passengers, *newPassenger)
	}
	return Passengers, nil
}

func (r *passengerRepository) FindByID(id string) (*entity.Passenger, error) {
	stmt, err := r.db.Prepare("SELECT * FROM passengers WHERE id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var model PassengerModel
	rows := stmt.QueryRow(id)
	rows.Scan(
		&model.ID,
		&model.Name,
		&model.Nickname,
		&model.RouteCode,
		&model.Goes,
		&model.Comesback,
		&model.RegisterConfirmed,
		&model.MonitorID,
		&model.active,
	)
	if err != nil {
		return nil, err
	}

	inputPassenger := factory.PassengerInputDTO{
		ID:                model.ID,
		Name:              model.Name,
		Nickname:          model.Nickname,
		RouteCode:         model.RouteCode,
		Goes:              model.Goes,
		Comesback:         model.Comesback,
		RegisterConfirmed: model.RegisterConfirmed,
		MonitorID:         model.MonitorID,
	}

	newPassenger, err := factory.PassengerFactory().Create(inputPassenger)
	if err != nil {
		return nil, err
	}

	newPassenger.ChangeGoNoGo(model.Goes, model.Comesback)
	return newPassenger, nil
}

func (r *passengerRepository) ListNotConfirmedPassengers() ([]entity.Passenger, error) {
	rows, err := r.db.Query("SELECT * FROM passengers WHERE register_confirmed = false  active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passengers []entity.Passenger
	for rows.Next() {
		var model PassengerModel
		err := rows.Scan(
			&model.ID,
			&model.Name,
			&model.Nickname,
			&model.RouteCode,
			&model.Goes,
			&model.Comesback,
			&model.MonitorID,
			&model.active,
		)
		if err != nil {
			return nil, err
		}

		inputPassenger := factory.PassengerInputDTO{
			ID:        model.ID,
			Name:      model.Name,
			Nickname:  model.Nickname,
			RouteCode: model.RouteCode,
			Goes:      model.Goes,
			Comesback: model.Comesback,
			MonitorID: model.MonitorID,
		}

		newPassenger, err := factory.PassengerFactory().Create(inputPassenger)
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

func (r *passengerRepository) FindByNameAndNickname(name, monitorID string) (*entity.Passenger, error) {
	stmt, err := r.db.Prepare("SELECT * FROM passengers WHERE name = ? and monitor_id = ? and active = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var model PassengerModel
	rows := stmt.QueryRow(name, monitorID)
	rows.Scan(
		&model.ID,
		&model.Name,
		&model.Nickname,
		&model.RouteCode,
		&model.Goes,
		&model.Comesback,
		&model.RegisterConfirmed,
		&model.MonitorID,
		&model.active,
	)
	if err != nil {
		return nil, err
	}

	inputPassenger := factory.PassengerInputDTO{
		ID:                model.ID,
		Name:              model.Name,
		Nickname:          model.Nickname,
		RouteCode:         model.RouteCode,
		Goes:              model.Goes,
		Comesback:         model.Comesback,
		RegisterConfirmed: model.RegisterConfirmed,
		MonitorID:         model.MonitorID,
	}

	newPassenger, err := factory.PassengerFactory().Create(inputPassenger)
	if err != nil {
		return nil, err
	}

	newPassenger.ChangeGoNoGo(model.Goes, model.Comesback)
	return newPassenger, nil
}

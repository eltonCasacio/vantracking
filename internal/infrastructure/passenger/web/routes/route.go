package routes

import (
	"database/sql"

	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/passenger/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/passenger/web/handlers"

	"github.com/go-chi/chi"
)

type passengerRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewPassengerRoutes(db *sql.DB, c *chi.Mux) *passengerRoutes {
	return &passengerRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *passengerRoutes) CreateRoutes() {
	repository := repo.NewPassengerRepository(dr.db)
	handler := handlers.NewMonitorHandler(repository)
	dr.chi.Route("/monitor", func(r chi.Router) {
		r.Get("/", handler.ListAll)
		r.Get("/not-confirmed", handler.ListNotConfirmed)
	})
}

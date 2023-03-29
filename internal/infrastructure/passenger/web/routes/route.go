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
	handler := handlers.NewPassengerHandler(repository)
	dr.chi.Route("/passenger", func(r chi.Router) {
		r.Post("/", handler.Register)
		r.Get("/", handler.ListAll)
		r.Get("/{id}", handler.Find)
		r.Put("/", handler.Update)
		r.Delete("/{id}", handler.Delete)
		r.Get("/not-confirmed/{routeCode}", handler.ListNotConfirmed)
		r.Put("/confirm-passenger", handler.ConfirmPassengerRegister)
		r.Get("/list-passenger/{monitor_id}", handler.ListByMonitorID)
		r.Get("/by-routecode/{routeCode}", handler.ListAllByRouteCode)
		r.Post("/finalize-route", handler.FinalizeRoute)
	})
}

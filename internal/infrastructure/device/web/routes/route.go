package routes

import (
	"database/sql"

	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/device/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/device/web/handlers"

	"github.com/go-chi/chi"
)

type deviceRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewDeviceRoutes(db *sql.DB, c *chi.Mux) *deviceRoutes {
	return &deviceRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *deviceRoutes) CreateRoutes() {
	repository := repo.NewDeviceRepository(dr.db)
	handler := handlers.NewDeviceHandler(repository)
	dr.chi.Route("/device", func(r chi.Router) {
		r.Post("/", handler.Register)
		//r.Get("/authenticate/{cpf}", handler.Authenticate)
		// r.Get("/", handler.ConsultAll)
		// r.Get("/{id}", handler.Consult)
		// r.Put("/", handler.Update)
		// r.Delete("/{id}", handler.Delete)
		// r.Post("/location", handler.SetLocation)
		// r.Post("/route", handler.CreateRoute)
		// r.Delete("/route/{code}", handler.DeleteRoute)
		// r.Get("/routes/{driverid}", handler.Routes)
	})
}

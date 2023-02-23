package routes

import (
	"database/sql"

	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/web/handlers"

	"github.com/go-chi/chi"
)

type driverRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewDriverRoutes(db *sql.DB, c *chi.Mux) *driverRoutes {
	return &driverRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *driverRoutes) CreateRoutes() {
	repository := repo.NewDriverRepository(dr.db)
	handler := handlers.NewDriverHandler(repository)
	dr.chi.Route("/driver", func(r chi.Router) {
		r.Post("/", handler.Register)
		r.Get("/", handler.ConsultAll)
		r.Get("/{id}", handler.Consult)
		r.Put("/", handler.Update)
		r.Delete("/{id}", handler.Delete)
		r.Post("/location", handler.SetLocation)
		r.Post("/route", handler.CreateRoute)
		r.Delete("/route/{id}", handler.DeleteRoute)
	})
}

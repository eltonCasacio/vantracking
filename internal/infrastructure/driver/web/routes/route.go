package routes

import (
	"database/sql"

	driverRepo "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/repository/mysql"
	driverHandler "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/web/handlers"

	"github.com/go-chi/chi"
)

type DriverRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewDriverRoutes(db *sql.DB, c *chi.Mux) *DriverRoutes {
	return &DriverRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *DriverRoutes) CreateRoutes() {
	driverRepository := driverRepo.NewDriverRepository(dr.db)
	driverHandler := driverHandler.NewDriverHandler(driverRepository)
	dr.chi.Route("/driver", func(r chi.Router) {
		r.Post("/", driverHandler.Register)
		r.Get("/", driverHandler.ConsultAll)
		r.Get("/{id}", driverHandler.Consult)
		r.Put("/", driverHandler.Update)
		r.Delete("/{id}", driverHandler.Delete)
	})
}

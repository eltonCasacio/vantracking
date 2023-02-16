package routes

import (
	"database/sql"

	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/monitor/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/monitor/web/handlers"

	"github.com/go-chi/chi"
)

type monitorRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewMonitorRoutes(db *sql.DB, c *chi.Mux) *monitorRoutes {
	return &monitorRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *monitorRoutes) CreateRoutes() {
	repository := repo.NewMonitorRepository(dr.db)
	monitorHandler := handlers.NewMonitorHandler(repository)
	dr.chi.Route("/monitor", func(r chi.Router) {
		r.Post("/", monitorHandler.Register)
		r.Get("/", monitorHandler.ConsultAll)
		r.Get("/{id}", monitorHandler.Consult)
		r.Put("/", monitorHandler.Update)
		r.Delete("/{id}", monitorHandler.Delete)
	})
}

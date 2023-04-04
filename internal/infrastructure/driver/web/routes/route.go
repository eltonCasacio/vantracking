package routes

import (
	"database/sql"

	"github.com/eltoncasacio/vantracking/configs"
	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/web/handlers"

	"github.com/go-chi/chi"
)

type driverRoutes struct {
	db     *sql.DB
	chi    *chi.Mux
	config *configs.Config
}

func NewDriverRoutes(db *sql.DB, c *chi.Mux, config *configs.Config) *driverRoutes {
	return &driverRoutes{
		db:     db,
		chi:    c,
		config: config,
	}
}

func (dr *driverRoutes) CreateRoutes() {
	repository := repo.NewDriverRepository(dr.db)
	handler := handlers.NewDriverHandler(repository, dr.config.TokenAuth, dr.config.JwtExperesIn)
	dr.chi.Route("/driver", func(r chi.Router) {
		r.Get("/authenticate/{cpf}", handler.Authenticate)
		r.Post("/", handler.Register)
		r.Get("/", handler.ConsultAll)
		r.Get("/{id}", handler.Consult)
		r.Put("/", handler.Update)
		r.Delete("/{id}", handler.Delete)
		r.Post("/location", handler.SetLocation)
		r.Post("/route", handler.CreateRoute)
		r.Delete("/route/{code}", handler.DeleteRoute)
		r.Get("/routes/{driverid}", handler.Routes)
	})
}

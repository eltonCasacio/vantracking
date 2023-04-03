package routes

import (
	"database/sql"

	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/partner/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/partner/web/handlers"

	"github.com/go-chi/chi"
)

type partnerRoutes struct {
	db  *sql.DB
	chi *chi.Mux
}

func NewPartnerRoutes(db *sql.DB, c *chi.Mux) *partnerRoutes {
	return &partnerRoutes{
		db:  db,
		chi: c,
	}
}

func (dr *partnerRoutes) CreateRoutes() {
	repository := repo.NewPartnerRepository(dr.db)
	handler := handlers.NewPartnerHandler(repository)
	dr.chi.Route("/partner", func(r chi.Router) {
		r.Post("/", handler.Register)
		r.Get("/", handler.FindAll)
		r.Get("/{id}", handler.FindByID)
		r.Get("/{city}", handler.FindByCity)
		r.Put("/", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}

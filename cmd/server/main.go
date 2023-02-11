package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/eltoncasacio/vantracking/configs"
	repo "github.com/eltoncasacio/vantracking/internal/infrastructure/repository/mysql"
	handlers "github.com/eltoncasacio/vantracking/internal/infrastructure/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBName,
		))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// driverLocations := map[string]string{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	driverRepository := repo.NewDriverRepository(db)
	driverHandler := handlers.NewDriverHandler(driverRepository)
	r.Route("/driver", func(r chi.Router) {
		r.Post("/", driverHandler.Register)
	})

	monitorRepository := repo.NewMonitorRepository(db)
	monitorHandler := handlers.NewMonitorHandler(monitorRepository)
	r.Route("/monitors", func(r chi.Router) {
		r.Post("/", monitorHandler.Register)
	})

	http.ListenAndServe(":8000", r)
}

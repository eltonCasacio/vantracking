package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/eltoncasacio/vantracking/configs"
	driverRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/web/routes"
	monitorRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/monitor/web/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	))
	if err != nil {
		panic(err)
	}
	// driverLocations := map[string]string{}

	chi := chi.NewRouter()
	chi.Use(middleware.Logger)
	chi.Use(middleware.Recoverer)

	driverRoutes.NewDriverRoutes(db, chi).CreateRoutes()
	monitorRoutes.NewMonitorRoutes(db, chi).CreateRoutes()

	http.ListenAndServe(":8000", chi)
}

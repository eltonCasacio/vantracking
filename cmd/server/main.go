package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/eltoncasacio/vantracking/configs"
	deviceRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/device/web/routes"
	driverRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/driver/web/routes"
	monitorRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/monitor/web/routes"
	partnerRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/partner/web/routes"
	passengerRoutes "github.com/eltoncasacio/vantracking/internal/infrastructure/passenger/web/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, config := Init()
	chi := chi.NewRouter()
	chi.Use(middleware.Logger)
	chi.Use(middleware.Recoverer)

	chi.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	driverRoutes.NewDriverRoutes(db, chi, config).CreateRoutes()
	monitorRoutes.NewMonitorRoutes(db, chi, config).CreateRoutes()
	passengerRoutes.NewPassengerRoutes(db, chi).CreateRoutes()
	partnerRoutes.NewPartnerRoutes(db, chi).CreateRoutes()
	deviceRoutes.NewDeviceRoutes(db, chi).CreateRoutes()

	http.ListenAndServe(fmt.Sprintf(":%v", config.WebServerPort), chi)
}

func Init() (*sql.DB, *configs.Config) {
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
	return db, config
}

package main

import (
	"fmt"

	fh "github.com/b92c/ineed-engine/internal/freelancers/handler"
	fs "github.com/b92c/ineed-engine/internal/freelancers/service"
	hch "github.com/b92c/ineed-engine/internal/healthcheck/handler"
	hcs "github.com/b92c/ineed-engine/internal/healthcheck/service"
	"github.com/b92c/ineed-engine/pkg/database"
	"github.com/b92c/ineed-engine/pkg/server"
)

func main() {
	// Database connection
	db, err := database.NewMySQLDB()
	if err != nil {
		panic(fmt.Sprintf("cannot connect to database: %s", err))
	}
	defer db.Close()

	// Dependency injection
	FreelanceSearchService := fs.NewFreelanceSearchService(db)
	freelanceSearchHandler := fh.NewFreelanceSearchHandler(FreelanceSearchService)

	HealthCheckService := hcs.NewHealthCheckService(db)
	HealthCheckHandler := hch.NewHealthCheckHandler(HealthCheckService)

	// Start server
	server := server.NewServer(freelanceSearchHandler, HealthCheckHandler)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

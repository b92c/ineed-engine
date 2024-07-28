package main

import (
	"fmt"

	c "github.com/b92c/ineed-engine/internal/adapters/controllers"
	uc "github.com/b92c/ineed-engine/internal/usecases"
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
	SearchUsecase := uc.NewSearchUsecase(db)
	searchController := c.NewSearchController(SearchUsecase)

	HealthCheckUsecase := uc.NewHealthCheckUsecase(db)
	HealthCheckController := c.NewHealthCheckController(HealthCheckUsecase)

	// Start server
	server := server.NewServer(searchController, HealthCheckController)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

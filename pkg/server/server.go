package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	c "github.com/b92c/ineed-engine/internal/adapters/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port                  int
	SearchController      *c.SearchController
	HealthCheckController *c.HealthCheckController
}

func NewServer(SearchController *c.SearchController, HealthCheckController *c.HealthCheckController) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	serverInstance := &Server{
		port:                  port,
		SearchController:      SearchController,
		HealthCheckController: HealthCheckController,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverInstance.port),
		Handler:      serverInstance.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Freelance routes
	r.GET("/search", s.SearchController.Search)

	// Health check routes
	r.GET("/health", s.HealthCheckController.HealthCheck)

	return r
}

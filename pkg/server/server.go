package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	fh "github.com/b92c/ineed-engine/internal/freelancers/handler"
	hch "github.com/b92c/ineed-engine/internal/healthcheck/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port             int
	freelanceHandler *fh.FreelanceSearchHandler
	healthHandler    *hch.HealthCheckHandler
}

func NewServer(freelanceHandler *fh.FreelanceSearchHandler, healthHandler *hch.HealthCheckHandler) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	serverInstance := &Server{
		port:             port,
		freelanceHandler: freelanceHandler,
		healthHandler:    healthHandler,
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
	r.GET("/freelance-search", s.freelanceHandler.HelloWorldHandler)

	// Health check routes
	r.GET("/health", s.healthHandler.HealthCheck)

	return r
}

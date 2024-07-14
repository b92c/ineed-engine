package handler

import (
	"net/http"

	"github.com/b92c/ineed-engine/internal/healthcheck/service"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	service *service.HealthCheckService
}

func NewHealthCheckHandler(service *service.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{service: service}
}

func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	status, err := h.service.GetHealthStatus(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

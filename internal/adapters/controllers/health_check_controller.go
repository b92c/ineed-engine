package controllers

import (
	"net/http"

	uc "github.com/b92c/ineed-engine/internal/usecases"
	"github.com/gin-gonic/gin"
)

type HealthCheckController struct {
	service *uc.HealthCheckUsecase
}

func NewHealthCheckController(service *uc.HealthCheckUsecase) *HealthCheckController {
	return &HealthCheckController{service: service}
}

func (h *HealthCheckController) HealthCheck(c *gin.Context) {
	status, err := h.service.GetHealthStatus(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

package handler

import (
	"net/http"

	"github.com/b92c/ineed-engine/internal/freelancers/service"
	"github.com/gin-gonic/gin"
)

type FreelanceSearchHandler struct {
	service *service.FreelanceSearchService
}

func NewFreelanceSearchHandler(service *service.FreelanceSearchService) *FreelanceSearchHandler {
	return &FreelanceSearchHandler{service: service}
}
func (f *FreelanceSearchHandler) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Jesus Cristo Salva!"

	c.JSON(http.StatusOK, resp)
}

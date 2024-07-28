package controllers

import (
	"fmt"
	"net/http"

	"github.com/b92c/ineed-engine/internal/usecases"
	"github.com/gin-gonic/gin"
)

type SearchController struct {
	service *usecases.SearchUsecase
}

func NewSearchController(service *usecases.SearchUsecase) *SearchController {
	return &SearchController{service: service}
}
func (f *SearchController) HelloWorldHandler(c *gin.Context) {
	professional := c.Query("professional")
	city := c.Query("city")

	fmt.Println("procurando pelo profissional: ", professional, " na cidade: ", city)

	resp := make(map[string]string)
	resp["message"] = "Jesus Cristo Salva!"

	c.JSON(http.StatusOK, resp)
}

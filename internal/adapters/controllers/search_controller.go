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
func (f *SearchController) Search(c *gin.Context) {
	professional := c.Query("professional")
	city := c.Query("city")

	fmt.Println("procurando pelo profissional: ", professional, " na cidade: ", city)

	r, err := f.service.Find(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Resultado do use case:", r)

	c.JSON(http.StatusOK, r)
}

package handlars

import (
	"mymachine707/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePerson(c *gin.Context) {
	var body models.CreatePersonModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
}

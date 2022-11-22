package handlars

import (
	"mymachine707/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handlars) RootCreatePerson(c *gin.Context) {
	var body models.CreatePersonModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := h.stg.AddPerson(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	person, err := h.stg.GetPersonByID(id.String())

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Person Created",
		Data:    person,
	})
}

func (h *Handlars) RootGetPersonByID(c *gin.Context) {

	id := c.Param("id")

	person, err := h.stg.GetPersonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Person successfuly get by id!",
		Data:    person,
	})
}

func (h *Handlars) RootUpdatePerson(c *gin.Context) {

	var body models.UpdatePersonModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.stg.UpdatePerson(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	person, err := h.stg.GetPersonByID(body.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Person successfuly UPDATED",
		Data:    person,
	})

}

func (h *Handlars) RootDeletePerson(c *gin.Context) {

	id := c.Param("id")

	person, err := h.stg.GetPersonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.stg.DeletePerson(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Person successfuly DELETED by id!",
		Data:    person,
	})
}

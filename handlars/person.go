package handlars

import (
	"mymachine707/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RootCreatePerson godoc
// @Summary     Creat Person
// @Description Creat a new person
// @Tags        Person
// @Accept      json
// @Produce     json
// @Param       person body     models.CreatePersonModul true "Person body"
// @Success     201    {object} models.JSONResult{data=models.Person}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/person [post]
func (h *Handlars) RootCreatePerson(c *gin.Context) {
	var body models.CreatePersonModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := h.Stg.AddPerson(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	person, err := h.Stg.GetPersonByID(id.String())

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

// RootGetPersonByID godoc
// @Summary     GetPersonyID
// @Description get an person by id
// @Tags        Person
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Person id"
// @Success     201 {object} models.JSONResult{data=models.Person}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/person/{id} [get]
func (h *Handlars) RootGetPersonByID(c *gin.Context) {

	id := c.Param("id")

	person, err := h.Stg.GetPersonByID(id)

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

// RootUpdatePerson godoc
// @Summary     Update Person
// @Description Update Person
// @Tags        Person
// @Accept      json
// @Produce     json
// @Param       person body     models.UpdatePersonModel true "Person body"
// @Success     201    {object} models.JSONResult{data=[]models.Person}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/person/ [put]
func (h *Handlars) RootUpdatePerson(c *gin.Context) {

	var body models.UpdatePersonModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.Stg.UpdatePerson(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	person, err := h.Stg.GetPersonByID(body.ID)

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

// RootDeletePerson ...
// @Summary     Delete Person
// @Description get person by id and delete this person
// @Tags        Person
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Person id"
// @Success     201 {object} models.JSONResult{data=models.Person}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/person/{id} [delete]
func (h *Handlars) RootDeletePerson(c *gin.Context) {

	id := c.Param("id")

	person, err := h.Stg.GetPersonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.DeletePerson(id)

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

// RootGetPersonList ...
// @Summary     List Person
// @Description GetPersonList
// @Tags        Person
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "100"
// @Param       search query    string false "search exapmle"
// @Success     200    {object} models.JSONResult{data=[]models.Person}
// @Router      /v1/person/ [get]
func (h *Handlars) RootGetPersonList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")
	search := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	personList, err := h.Stg.GetPersonList(offset, limit, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "PersonList successfuly get!",
		Data:    personList,
	})
}

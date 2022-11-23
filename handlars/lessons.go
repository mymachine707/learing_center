package handlars

import (
	"fmt"
	"mymachine707/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RootCreateLesson godoc
// @Summary     Creat Lesson
// @Description Creat a new person
// @Tags        Lesson
// @Accept      json
// @Produce     json
// @Param       person body     models.CreateLessonModels true "Lesson body"
// @Success     201    {object} models.JSONResult{data=models.Lesson}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/lesson [post]
func (h *Handlars) RootCreateLesson(c *gin.Context) {
	var body models.CreateLessonModels

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	id := uuid.New()
	err := h.Stg.AddLesson(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := h.Stg.GetLessonByID(id.String())

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Lesson successfuly created!",
		Data:    result,
	})
	return

}

// RootGetLessonByID godoc
// @Summary     Get Lesson By ID
// @Description Creat a new person
// @Tags        Lesson
// @Accept      json
// @Produce     json
// @Param       id path     string true "Lesson body"
// @Success     201    {object} models.JSONResult{data=models.Lesson}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/lesson [get]
func (h *Handlars) RootGetLessonByID(c *gin.Context) {

	id := c.Param("id")

	fmt.Printf("this id is : %s", id)

	lesson, err := h.Stg.GetLessonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Lesson successfuly get by Id!",
		Data:    lesson,
	})
}

// RootGetLessonList ...
// @Summary     List Lesson
// @Description GetLessonList
// @Tags        Lesson
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "100"
// @Param       search query    string false "search exapmle"
// @Success     200    {object} models.JSONResult{data=[]models.Lesson}
// @Router      /v1/lesson/ [get]
func (h *Handlars) RootGetLessonList(c *gin.Context) {
	offsetStr := c.DefaultQuery("Offset", "0")
	limitStr := c.DefaultQuery("Limit", "100")
	search := c.DefaultQuery("search", "100")

	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	listLesson, err := h.Stg.GetListLesson(offset, limit, search)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "List Lesson successfuly get by Id!",
		Data:    listLesson,
	})
}

// RootUpdateLesson godoc
// @Summary     Update Lesson
// @Description Update Lesson
// @Tags        Lesson
// @Accept      json
// @Produce     json
// @Param       lesson body     models.UpdateLessonModels true "Lesson body"
// @Success     201    {object} models.JSONResult{data=[]models.Lesson}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v1/lesson/ [put]
func (h *Handlars) RootUpdateLesson(c *gin.Context) {
	var body models.UpdateLessonModels

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateLessonByID(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.Stg.GetLessonByID(body.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Lesson successfuly UPDATED!",
		Data:    resp,
	})
}

// RootDeleteLesson ...
// @Summary     Delete Lesson
// @Description get lesson by id and delete this lesson
// @Tags        Lesson
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Lesson id"
// @Success     201 {object} models.JSONResult{data=models.Lesson}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/lesson/{id} [delete]
func (h *Handlars) RootDeleteLesson(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.Stg.GetLessonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.DeleteLessonByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Lesson successfuly Deleted!",
		Data:    resp,
	})
}

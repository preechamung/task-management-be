package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

type AddProjectRequestBody struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) AddProject(c *gin.Context) {
	body := AddProjectRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var project models.Project

	project.Name = body.Name
	project.Title = body.Title
	project.Description = body.Description

	if result := h.DB.Create(&project); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &project)
}

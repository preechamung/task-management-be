package project_statuses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

type AddProjectStatusesRequestBody struct {
	Name      string `json:"name"`
	ProjectId uint   `json:"project_id"`
	Order     uint   `json:"order"`
}

func (h handler) AddProjectStatus(c *gin.Context) {
	body := AddProjectStatusesRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var project_status models.ProjectStatus

	project_status.Name = body.Name
	project_status.ProjectId = body.ProjectId
	project_status.Order = body.Order

	if result := h.DB.Create(&project_status); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &project_status)
}

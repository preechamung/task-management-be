package project_statuses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

func (h handler) DeleteProjectStatus(c *gin.Context) {
	id := c.Param("id")

	var status models.ProjectStatus

	if result := h.DB.First(&status, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&status)

	c.Status(http.StatusOK)
}

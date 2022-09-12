package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

func (h handler) GetProjects(c *gin.Context) {
	var projects []models.Project

	if result := h.DB.Preload("Permissions").Preload("Statuses").Find(&projects); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &projects)
}

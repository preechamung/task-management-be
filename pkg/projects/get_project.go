package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

func (h handler) GetProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project

	if result := h.DB.Preload("Permissions").Preload("Statuses").Find(&project, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &project)
}

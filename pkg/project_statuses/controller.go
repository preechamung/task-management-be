package project_statuses

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/project_statuses")
	routes.POST("/", h.AddProjectStatus)
	routes.DELETE("/:id", h.DeleteProjectStatus)
}

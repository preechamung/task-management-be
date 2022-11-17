package oauth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Route(rg *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := rg.Group("/oauth")
	routes.GET("/google", h.GoogleOAuth)
}

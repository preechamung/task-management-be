package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-be/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Route(rg *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := rg.Group("/auth")
	routes.POST("/register", h.SignUpUser)
	routes.POST("/login", h.SignInUser)
	routes.POST("/refresh", h.RefreshAccessToken)
	routes.GET("/logout", middleware.DeserializeUser(h.DB), h.LogoutUser)
}

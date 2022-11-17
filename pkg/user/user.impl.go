package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
	"github.com/preechamung/task-management-fe/utils"
)

func (h handler) PostUser(c *gin.Context) {
	body := &models.SignUpInput{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	hashedPassword := utils.HashPassword(body.Password)

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPassword,
		Provider: body.Provider,
		Photo:    body.Photo,
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": result.Error})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

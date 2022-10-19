package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
	"github.com/preechamung/task-management-fe/utils"
)

func (h handler) SignUpUser(c *gin.Context) {
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
		Provider: "password",
		Photo:    body.Photo,
	}
	result := h.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": result.Error})
		return
	}

	response := models.FilteredResponse(&user)

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": response})
}

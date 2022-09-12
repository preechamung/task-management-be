package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/models"
)

type AddUserRequestBody struct {
	Name string `json:"name"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Name = body.Name

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

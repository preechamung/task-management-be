package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-be/pkg/common/config"
	"github.com/preechamung/task-management-be/pkg/common/models"
	"github.com/preechamung/task-management-be/utils"
)

func (h handler) SignUpUser(c *gin.Context) {
	var body *models.SignUpInput

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
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": result.Error})
		return
	}

	response := models.FilteredResponse(&user)

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": response})
}

func (h handler) SignInUser(c *gin.Context) {
	var credentials *models.SignInInput

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := h.DB.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
		return
	}

	if err := utils.ComparePassword(user.Password, credentials.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
		return
	}

	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.Id, config.AccessTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.Id, config.RefreshTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (h handler) RefreshAccessToken(c *gin.Context) {
	message := "could not refresh access token"

	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := config.LoadConfig()

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := h.DB.First(&user, "id = ?", sub)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.Id, config.AccessTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (h handler) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

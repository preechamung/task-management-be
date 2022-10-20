package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/config"
	"github.com/preechamung/task-management-fe/pkg/common/models"
	"github.com/preechamung/task-management-fe/utils"
	"gorm.io/gorm/clause"
)

func (h handler) GoogleOAuth(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	googleUser, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	user := models.User{
		Name:     googleUser.Name,
		Email:    googleUser.Email,
		Provider: "google",
		Verified: true,
		Photo:    googleUser.Picture,
	}

	// h.DB.Create(&user)
	result := h.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "email"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"name":     user.Name,
			"email":    user.Email,
			"provider": user.Provider,
			"verified": user.Verified,
			"photo":    user.Photo,
		}),
	}).Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": result.Error})
		return
	}

	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.Id, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.Id, config.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(config.Origin, pathUrl))
}

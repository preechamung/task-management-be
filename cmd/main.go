package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/auth"
	"github.com/preechamung/task-management-fe/pkg/common/config"
	"github.com/preechamung/task-management-fe/pkg/common/db"
	"github.com/preechamung/task-management-fe/pkg/oauth"
	"github.com/preechamung/task-management-fe/pkg/user"
	"github.com/spf13/viper"
)

var (
	server *gin.Engine
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	config, _ := config.LoadConfig()

	fmt.Println(config)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true
	server = gin.Default()

	server.Use(cors.New(corsConfig))

	// connedt postgrest
	h := db.Init(config.PostgresSource)

	router := server.Group("/api")

	// register more routes here
	user.Route(router, h)
	auth.Route(router, h)
	oauth.Route(router, h)

	// not found route
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.Port))
}

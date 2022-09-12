package main

import (
	"github.com/gin-gonic/gin"
	"github.com/preechamung/task-management-fe/pkg/common/db"
	"github.com/preechamung/task-management-fe/pkg/project_statuses"
	"github.com/preechamung/task-management-fe/pkg/projects"
	"github.com/preechamung/task-management-fe/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	users.RegisterRoutes(r, h)
	projects.RegisterRoutes(r, h)
	project_statuses.RegisterRoutes(r, h)

	// register more routes here

	r.Run(port)
}

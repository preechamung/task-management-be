package db

import (
	"log"

	"github.com/preechamung/task-management-fe/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.ProjectStatus{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.User{})

	return db
}

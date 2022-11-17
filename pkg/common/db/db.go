package db

import (
	"fmt"
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

	db.AutoMigrate(&models.User{})

	fmt.Println("PostgreSQL connected successfully...")

	return db
}

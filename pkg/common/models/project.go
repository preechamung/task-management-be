package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id          uint            `json:"id" gorm:"primaryKey"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Statuses    []ProjectStatus `json:"statuses" gorm:"foreignKey:ProjectId"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"deleted_at" gorm:"index"`
}

type ProjectStatus struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	ProjectId uint           `json:"project_id"`
	Order     uint           `json:"order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

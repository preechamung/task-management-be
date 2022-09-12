package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Permissions []Permission    `gorm:"foreignKey:ProjectId" json:"permissions"`
	Statuses    []ProjectStatus `gorm:"foreignKey:ProjectId" json:"statuses"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type ProjectStatus struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	ProjectId uint           `json:"project_id"`
	Order     uint           `json:"order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Permission struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Role      string `json:"role"`
	ProjectId uint   `json:"project_id"`
	UserId    uint   `json:"user_id"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Provider string `json:"provider"`
	Photo    string `json:"photo,omitempty"`
	Verified bool   `json:"verified"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// db model
type User struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	Provider  string         `json:"provider"`
	Photo     string         `json:"photo"`
	Verified  bool           `json:"verified"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Provider  string    `json:"provider"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		Provider:  user.Provider,
		Photo:     user.Photo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

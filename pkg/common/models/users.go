package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
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
	Id        uuid.UUID      `gorm:"primaryKey" json:"id"`
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
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Provider  string    `json:"provider"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilteredResponse(u *User) UserResponse {
	return UserResponse{
		Id:        u.Id,
		Email:     u.Email,
		Name:      u.Name,
		Provider:  u.Provider,
		Photo:     u.Photo,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV4()

	u.Id = uuid
	if err != nil {
		err = errors.New("can't save invalid data")
	}
	return
}

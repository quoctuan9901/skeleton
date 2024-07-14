package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model người dùng
type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Password  string         `json:"-" validate:"required,min=6"`
	Level     int            `gorm:"default:1" json:"level"`
	Status    int            `gorm:"default:1" json:"status"`
	FullName  string         `json:"full_name"`
	Token     string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// RegisterInput input cho đăng ký
type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name"`
}

// LoginInput input cho đăng nhập
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

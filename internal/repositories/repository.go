package repositories

import (
	"github.com/google/uuid"
	"skeleton/internal/models"
)

// UserRepository interface cho repository user
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(ID uuid.UUID) error
}

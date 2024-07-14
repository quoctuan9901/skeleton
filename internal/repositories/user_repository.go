package repositories

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"skeleton/internal/models"
)

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository tạo user repository mới
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// CreateUser tạo người dùng mới
func (r *userRepository) CreateUser(user *models.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return fmt.Errorf("lỗi khi tạo người dùng: %w", err)
	}
	return nil
}

// GetUserByEmail lấy thông tin người dùng theo email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("lỗi khi lấy thông tin người dùng: %w", err)
	}
	return &user, nil
}

// GetUserByID lấy thông tin người dùng theo ID
func (r *userRepository) GetUserByID(ID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser cập nhật thông tin người dùng
func (r *userRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

// DeleteUser xóa người dùng
func (r *userRepository) DeleteUser(ID uuid.UUID) error {
	return r.DB.Delete(&models.User{}, "id = ?", ID).Error
}

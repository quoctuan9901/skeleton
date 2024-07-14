package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"skeleton/internal/models"
	"skeleton/internal/repositories"
	"skeleton/internal/utils"
)

// UserService service cho user
type UserService struct {
	userRepository repositories.UserRepository
	logger         *zap.Logger
}

// NewUserService tạo UserService mới
func NewUserService(userRepository repositories.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

// CreateUser tạo người dùng mới
func (s *UserService) CreateUser(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := s.userRepository.CreateUser(user); err != nil {
		return fmt.Errorf("lỗi khi tạo người dùng: %w", err)
	}
	return nil
}

// GetUserByEmail lấy thông tin người dùng theo email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

// GetUserByID lấy thông tin người dùng theo ID
func (s *UserService) GetUserByID(ID uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		s.logger.Error("Lỗi get user by ID", zap.String("id", ID.String()), zap.Error(err))
		return nil, utils.ErrInternalServerError
	}
	return user, nil
}

// UpdateUser cập nhật thông tin người dùng
func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.userRepository.UpdateUser(user); err != nil {
		return fmt.Errorf("lỗi khi cập nhật người dùng: %w", err)
	}
	return nil
}

// DeleteUser xóa người dùng
func (s *UserService) DeleteUser(ID uuid.UUID) error {
	if err := s.userRepository.DeleteUser(ID); err != nil {
		return fmt.Errorf("lỗi khi xóa người dùng: %w", err)
	}
	return nil
}

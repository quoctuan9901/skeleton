package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"skeleton/internal/models"
	"skeleton/internal/repositories"
	"skeleton/internal/utils"
)

// AuthService service cho authentication
type AuthService struct {
	userRepository repositories.UserRepository
	logger         *zap.Logger
}

// NewAuthService tạo AuthService mới
func NewAuthService(userRepository repositories.UserRepository, logger *zap.Logger) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		logger:         logger,
	}
}

// Register đăng ký người dùng mới
func (s *AuthService) Register(input *models.RegisterInput) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := &models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		FullName: input.FullName,
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, utils.ErrConflict
		}
		return nil, fmt.Errorf("lỗi khi tạo người dùng: %w", err)
	}

	return user, nil
}

// Login đăng nhập cho người dùng
func (s *AuthService) Login(input *models.LoginInput) (string, error) {
	user, err := s.userRepository.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", utils.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", utils.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

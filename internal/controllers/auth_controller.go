package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"skeleton/internal/models"
	"skeleton/internal/services"
	"skeleton/internal/utils"
)

// AuthController controller cho authentication
type AuthController struct {
	authService *services.AuthService
	logger      *zap.Logger
}

// NewAuthController tạo AuthController mới
func NewAuthController(authService *services.AuthService, logger *zap.Logger) *AuthController {
	return &AuthController{
		authService: authService,
		logger:      logger,
	}
}

// Register xử lý đăng ký người dùng mới
func (ctrl *AuthController) Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.logger.Error("Lỗi khi bind JSON", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		ctrl.logger.Error("Lỗi validation", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	user, err := ctrl.authService.Register(&input)
	if err != nil {
		ctrl.logger.Error("Lỗi khi đăng ký người dùng", zap.Error(err))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusCreated, "Tạo người dùng thành công", user)
}

// Login xử lý đăng nhập cho người dùng
func (ctrl *AuthController) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.logger.Error("Lỗi khi bind JSON", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		ctrl.logger.Error("Lỗi validation", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	token, err := ctrl.authService.Login(&input)
	if err != nil {
		ctrl.logger.Error("Lỗi khi đăng nhập", zap.Error(err))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusOK, "Đăng nhập thành công", gin.H{"token": token})
}

package controllers

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"skeleton/internal/models"
	"skeleton/internal/services"
	"skeleton/internal/utils"
)

// UserController controller cho user
type UserController struct {
	userService *services.UserService
	logger      *zap.Logger
}

// NewUserController tạo UserController mới
func NewUserController(userService *services.UserService, logger *zap.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

// CreateUser tạo người dùng mới
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.logger.Error("Lỗi khi bind JSON", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctrl.logger.Error("Lỗi validation", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		ctrl.logger.Error("Lỗi khi tạo người dùng", zap.Error(err))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusCreated, "Tạo người dùng thành công", user)
}

// GetUser lấy thông tin người dùng
func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Error("Lỗi khi parse ID", zap.Error(err), zap.String("id", id))
		utils.ApiResponseError(c, err)
		return
	}

	user, err := ctrl.userService.GetUserByID(ID)
	if err != nil {
		ctrl.logger.Error("Lỗi khi lấy người dùng", zap.Error(err), zap.String("userId", id))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusOK, "Lấy thông tin người dùng thành công", user)
}

// UpdateUser cập nhật thông tin người dùng
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Error("Lỗi khi parse ID", zap.Error(err), zap.String("id", id))
		utils.ApiResponseError(c, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.logger.Error("Lỗi khi bind JSON", zap.Error(err))
		utils.HandleValidationError(c, err)
		return
	}

	user.ID = ID

	if err := ctrl.userService.UpdateUser(&user); err != nil {
		ctrl.logger.Error("Lỗi khi cập nhật người dùng", zap.Error(err), zap.String("userId", id))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusOK, "Cập nhật người dùng thành công", user)
}

// DeleteUser xóa người dùng
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Error("Lỗi khi parse ID", zap.Error(err), zap.String("id", id))
		utils.ApiResponseError(c, err)
		return
	}

	if err := ctrl.userService.DeleteUser(ID); err != nil {
		ctrl.logger.Error("Lỗi khi xóa người dùng", zap.Error(err), zap.String("userId", id))
		utils.ApiResponseError(c, err)
		return
	}

	utils.ApiResponseSuccess(c, http.StatusOK, "Xóa người dùng thành công", nil)
}

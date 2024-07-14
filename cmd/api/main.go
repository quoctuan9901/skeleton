package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"skeleton/internal/config"
	"skeleton/internal/controllers"
	"skeleton/internal/http"
	"skeleton/internal/http/middleware"
	"skeleton/internal/models"
	"skeleton/internal/repositories"
	"skeleton/internal/services"
	"skeleton/pkg/mylogger"
	"skeleton/platform/database"
)

func main() {
	// Load biến môi trường
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Lỗi khi tải file .env:", err)
	}

	// Khởi tạo config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Lỗi khi khởi tạo config: %v", err)
	}

	// Khởi tạo logger
	logger, err := mylogger.NewLogger("error.log")
	if err != nil {
		log.Fatalf("Lỗi khi tạo logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// Khởi tạo database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		logger.Fatal("Lỗi khi kết nối database", zap.Error(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Lỗi khi đóng kết nối database", zap.Error(err))
		}
	}()

	// Migration
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		logger.Fatal("Lỗi khi chạy migration", zap.Error(err))
	}

	// Khởi tạo repository
	userRepository := repositories.NewUserRepository(db.DB)

	// Khởi tạo service
	userService := services.NewUserService(userRepository, logger)
	authService := services.NewAuthService(userRepository, logger)

	// Khởi tạo controller
	userController := controllers.NewUserController(userService, logger)
	authController := controllers.NewAuthController(authService, logger)

	// Khởi tạo Gin router
	router := gin.Default()

	// Sử dụng middleware log request
	router.Use(middleware.LoggingMiddleware(logger))

	// Setup các routes
	http.SetupRoutes(router, userController, authController)

	// Chạy server
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logger.Fatal("Lỗi khi chạy server", zap.Error(err))
	}
}

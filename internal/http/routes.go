package http

import (
	"github.com/gin-gonic/gin"
	"skeleton/internal/controllers"
	"skeleton/internal/http/middleware"
)

// SetupRoutes thiết lập các routes cho ứng dụng
func SetupRoutes(router *gin.Engine, userController *controllers.UserController, authController *controllers.AuthController) {
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/users/:id", userController.GetUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}
}

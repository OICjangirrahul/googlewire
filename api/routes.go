package api

import (
	"googlewire/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", userController.CreateUser) // Create user
		userGroup.GET("/:id", userController.GetUser) // Get user by ID
	}
}

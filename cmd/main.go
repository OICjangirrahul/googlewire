package main

import (
	"log"

	"googlewire/api"
	"googlewire/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Setup routes and controllers using Wire
	userController, err := internal.InitializeUserController() // Initialize via Wire
	if err != nil {
		log.Fatal("Failed to initialize user controller:", err)
	}
	api.SetupRoutes(r, userController)

	// Start the server
	if err := r.Run(":4000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

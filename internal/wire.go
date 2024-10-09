//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"googlewire/internal/controllers"
	"googlewire/internal/models" // Import your models package
	"googlewire/internal/repositories"
	"googlewire/internal/services"
)

// InitializeUserController initializes the UserController and its dependencies.
func InitializeUserController() (*controllers.UserController, error) {
	wire.Build(
		NewDatabaseConnection,          // Build a database connection
		repositories.NewUserRepository, // Build a user repository
		services.NewUserService,        // Build a user service
		controllers.NewUserController,  // Build a user controller
	)
	return &controllers.UserController{}, nil // This line is never reached; Wire will generate the code
}

// NewDatabaseConnection creates a new GORM database connection and performs auto-migration.
func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "root:password@tcp(db:3306)/first?charset=utf8mb4&parseTime=True&loc=Local" // Update with your MySQL config
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // Return an error if database connection fails
	}

	// Perform auto-migration for your models
	if err := db.AutoMigrate(&models.User{}); err != nil { // Replace with your actual models
		return nil, err // Return an error if migration fails
	}

	return db, nil
}

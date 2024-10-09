package services

import (
	"googlewire/internal/models"
	"googlewire/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser calls the repository to save the user
func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// GetUser calls the repository to retrieve a user
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUser(id)
}

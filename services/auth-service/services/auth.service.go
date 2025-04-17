package services

import (
	"hauparty/libs/common/models"
	"hauparty/services/auth-service/repository"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type authService struct {
	repo repository.UserRepository
}

// Login implements AuthService.
func (a *authService) Login(email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// Register implements AuthService.
func (a *authService) Register(user *models.User) error {
	panic("unimplemented")
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

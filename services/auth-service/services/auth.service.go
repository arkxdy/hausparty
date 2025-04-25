package services

import (
	models "hausparty/libs/db/models/users"
	"hausparty/libs/db/repository"
)

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type authService struct {
	repo repository.IUserRepository
}

// Login implements AuthService.
func (a *authService) Login(email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// Register implements AuthService.
func (a *authService) Register(user *models.User) error {
	panic("unimplemented")
}

func NewAuthService(repo repository.IUserRepository) AuthService {
	return &authService{repo: repo}
}

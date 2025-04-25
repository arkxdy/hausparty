package services

import (
	models "hausparty/libs/db/models/users"
	"hausparty/libs/db/repository"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

// Login implements UserService.
func (u *userService) Login(email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// Register implements UserService.
func (u *userService) Register(user *models.User) error {
	panic("unimplemented")
}

func NewUserService(repo repository.IUserRepository) UserService {
	return &userService{repo: repo}
}

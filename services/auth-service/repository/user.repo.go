package repository

import "hauparty/libs/common/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
}

type userRepository struct {
	dbUrl string
}

// FindByEmail implements UserRepository.
func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// Save implements UserRepository.
func (u *userRepository) Save(user *models.User) error {
	panic("unimplemented")
}

func NewUserRepository(dbUrl string) UserRepository {
	return &userRepository{dbUrl: dbUrl}
}

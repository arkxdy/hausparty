package repository

import (
	"context"
	models "hausparty/libs/db/models/users"
	"time"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUserProfile(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

// GORM Implementation
type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return &user, err
}

func (r *userRepository) UpdateUserProfile(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Multiple operations in transaction
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return tx.Model(&models.AuthCredentials{}).
			Where("user_id = ?", user.ID).
			Update("updated_at", time.Now()).Error
	})
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

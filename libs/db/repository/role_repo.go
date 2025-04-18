package repository

import (
	"context"
	models "hauparty/libs/db/models/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	AssignRole(ctx context.Context, userID string, roleName string) error
	RemoveRole(ctx context.Context, userID string, roleName string) error
	HasRole(ctx context.Context, userID string, roleName string) (bool, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) AssignRole(ctx context.Context, userID, roleName string) error {
	var role models.Role
	err := r.db.WithContext(ctx).First(&role, "name = ?", roleName).Error
	if err != nil {
		return err
	}

	userRole := models.UserRole{
		UserID: uuid.MustParse(userID),
		RoleID: role.ID,
	}
	return r.db.WithContext(ctx).Create(&userRole).Error
}

func (r *roleRepository) RemoveRole(ctx context.Context, userID, roleName string) error {
	var role models.Role
	err := r.db.WithContext(ctx).First(&role, "name = ?", roleName).Error
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", userID, role.ID).
		Delete(&models.UserRole{}).Error
}

func (r *roleRepository) HasRole(ctx context.Context, userID, roleName string) (bool, error) {
	var role models.Role
	err := r.db.WithContext(ctx).First(&role, "name = ?", roleName).Error
	if err != nil {
		return false, err
	}

	var count int64
	err = r.db.WithContext(ctx).Model(&models.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, role.ID).
		Count(&count).Error
	return count > 0, err
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &roleRepository{db: db}
}

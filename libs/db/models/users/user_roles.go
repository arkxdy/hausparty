package models

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID uuid.UUID `gorm:"type:uuid;primaryKey"`
}

// user_id         UUID REFERENCES users(user_id) ON DELETE CASCADE,
// role_id         INT REFERENCES roles(role_id) ON DELETE CASCADE,
// PRIMARY KEY (user_id, role_id)

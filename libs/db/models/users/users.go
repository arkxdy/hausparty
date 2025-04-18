package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email         string    `gorm:"uniqueIndex;not null"`
	Username      string    `gorm:"uniqueIndex"`
	ProfilePicURL string
	IsVerified    bool `gorm:"default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	// Relationships
	AuthCredentials AuthCredentials `gorm:"foreignKey:UserID"`
	Roles           []Role          `gorm:"many2many:user_roles;"`
	Sessions        []Session
	AdminActions    []AdminAction `gorm:"foreignKey:AdminID"`
}

// user_id         UUID PRIMARY KEY,
// email           VARCHAR(255) UNIQUE NOT NULL,
// username        VARCHAR(50) UNIQUE,
// profile_pic_url VARCHAR(255),        -- S3/Cloud Storage URL
// created_at      TIMESTAMPTZ DEFAULT NOW(),
// updated_at      TIMESTAMPTZ,
// is_verified     BOOLEAN DEFAULT FALSE

package models

import "github.com/google/uuid"

type AuthCredentials struct {
	UserID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PasswordHash  string    `gorm:"type:varchar(255)"`                         // Nullable for OAuth
	OAuthProvider string    `gorm:"type:varchar(50);index:oauth_provider_id"`  // Nullable for email
	OAuthID       string    `gorm:"type:varchar(255);index:oauth_provider_id"` // Nullable for email
	Email         string    `gorm:"type:varchar(255);uniqueIndex"`             // Add email field
}

// user_id         UUID PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
// password_hash   VARCHAR(255),        -- Nullable for OAuth users
// oauth_provider  VARCHAR(50),         -- e.g., 'google', 'facebook'
// oauth_id        VARCHAR(255)         -- Unique ID from OAuth provider

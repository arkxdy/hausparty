package models

import "github.com/google/uuid"

type AuthCredentials struct {
	UserID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PasswordHash  string
	OAuthProvider string
	OAuthID       string `gorm:"index"`
}

// user_id         UUID PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
// password_hash   VARCHAR(255),        -- Nullable for OAuth users
// oauth_provider  VARCHAR(50),         -- e.g., 'google', 'facebook'
// oauth_id        VARCHAR(255)         -- Unique ID from OAuth provider

package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        string    `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

// session_id      UUID PRIMARY KEY,
// user_id         UUID REFERENCES users(user_id) ON DELETE CASCADE,
// token           TEXT NOT NULL,
// expires_at      TIMESTAMPTZ NOT NULL

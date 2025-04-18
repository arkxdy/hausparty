package models

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `gorm:"uniqueIndex;not null"`
}

// role_id         SERIAL PRIMARY KEY,
// name            VARCHAR(20) UNIQUE NOT NULL  -- e.g., 'user', 'host', 'admin'

package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendee struct {
	PartyID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Status   string    `gorm:"default:'pending'"`
	JoinedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// attendee_id     SERIAL PRIMARY KEY,
// party_id        UUID REFERENCES parties(party_id) ON DELETE CASCADE,
// user_id         UUID NOT NULL,                -- References users.user_id (cross-DB)
// status          VARCHAR(20) DEFAULT 'pending',-- e.g., 'pending', 'confirmed'
// joined_at       TIMESTAMPTZ DEFAULT NOW()

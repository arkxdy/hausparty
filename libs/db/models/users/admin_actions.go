package models

import (
	"time"

	"github.com/google/uuid"
)

type AdminAction struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	AdminID       uuid.UUID `gorm:"type:uuid;index"`
	ActionType    string    `gorm:"not null"`
	TargetUserID  uuid.UUID `gorm:"type:uuid;index"`
	TargetPartyID uuid.UUID `gorm:"type:uuid;index"`
	Details       string
	CreatedAt     time.Time
}

// action_id       SERIAL PRIMARY KEY,
// admin_id        UUID REFERENCES users(user_id) ON DELETE SET NULL,
// action_type     VARCHAR(50) NOT NULL,        -- e.g., 'delete_party', 'ban_user'
// target_user_id  UUID REFERENCES users(user_id) ON DELETE SET NULL,
// target_party_id UUID,                        -- No FK (since parties are in another DB)
// details         TEXT,
// created_at      TIMESTAMPTZ DEFAULT NOW()

package models

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ReporterID      uuid.UUID `gorm:"type:uuid;index;not null"` // References User.ID
	ReportedPartyID uuid.UUID `gorm:"type:uuid;index"`          // Nullable
	ReportedUserID  uuid.UUID `gorm:"type:uuid;index"`          // Nullable
	Reason          string    `gorm:"not null"`
	Status          string    `gorm:"default:'pending'"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// report_id       SERIAL PRIMARY KEY,
// reporter_id     UUID NOT NULL,                -- References users.user_id (cross-DB)
// reported_party_id UUID REFERENCES parties(party_id) ON DELETE CASCADE,
// reported_user_id UUID,                        -- References users.user_id (cross-DB)
// reason          TEXT NOT NULL,
// status          VARCHAR(20) DEFAULT 'pending',-- e.g., 'pending', 'resolved'
// created_at      TIMESTAMPTZ DEFAULT NOW()

package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/twpayne/go-geom/encoding/ewkb"
)

type Party struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	HostID      uuid.UUID `gorm:"type:uuid;index;not null"` // References User.ID in auth DB
	Title       string    `gorm:"not null"`
	Description string
	Location    ewkb.Point `gorm:"not null"` // Store as text (e.g., "lat,lng")
	Address     string     `gorm:"not null"`
	StartTime   time.Time  `gorm:"index"`
	EndTime     time.Time
	Capacity    int
	IsPrivate   bool           `gorm:"default:false"`
	InviteCode  string         `gorm:"index"`
	Status      string         `gorm:"default:'active'"`
	GalleryURLs pq.StringArray `gorm:"type:text[]"`

	// Relationships
	Attendees []Attendee
	Reports   []Report `gorm:"foreignKey:ReportedPartyID;references:ID;constraint:OnDelete:CASCADE"`
}

// party_id        UUID PRIMARY KEY,
// host_id         UUID NOT NULL,                -- References users.user_id (cross-DB)
// title           VARCHAR(100) NOT NULL,
// description     TEXT,
// location        GEOGRAPHY(Point, 4326),       -- PostGIS point (lat/long)
// address         VARCHAR(255),                 -- Human-readable address
// start_time      TIMESTAMPTZ NOT NULL,
// end_time        TIMESTAMPTZ,
// capacity        INT,
// is_private      BOOLEAN DEFAULT FALSE,
// invite_code     VARCHAR(10),                  -- For private parties
// status          VARCHAR(20) DEFAULT 'active'  -- e.g., 'active', 'cancelled'

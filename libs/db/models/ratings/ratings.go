package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// For MongoDB (using mongo-driver instead of GORM)
type Rating struct {
	ID           primitive.ObjectID `bson:"_id"`
	PartyID      uuid.UUID          `bson:"party_id"`
	UserID       uuid.UUID          `bson:"user_id"`
	Rating       int                `bson:"rating"`
	Comment      string             `bson:"comment"`
	HostResponse string             `bson:"host_response"`
	CreatedAt    time.Time          `bson:"created_at"`
}

// {
// 	"rating_id": ObjectId,
// 	"party_id": "UUID (from Postgres)",
// 	"user_id": "UUID (from Postgres)",
// 	"rating": 1-5,
// 	"comment": "text",
// 	"host_response": "text (optional)",
// 	"created_at": ISODate
//   }

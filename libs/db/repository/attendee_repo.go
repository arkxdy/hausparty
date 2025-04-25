package repository

import (
	"context"
	models "hausparty/libs/db/models/party"

	"gorm.io/gorm"
)

type IAttendeeRepository interface {
	AddAttendee(ctx context.Context, attendee *models.Attendee) error
	UpdateAttendeeStatus(ctx context.Context, partyID, userID string, status string) error
	GetAttendees(ctx context.Context, partyID string) ([]models.Attendee, error)
	RemoveAttendee(ctx context.Context, partyID, userID string) error
}

type attendeeRepository struct {
	db *gorm.DB
}

// AddAttendee implements IAttendeeRepository.
func (a *attendeeRepository) AddAttendee(ctx context.Context, attendee *models.Attendee) error {
	panic("unimplemented")
}

// GetAttendees implements IAttendeeRepository.
func (a *attendeeRepository) GetAttendees(ctx context.Context, partyID string) ([]models.Attendee, error) {
	panic("unimplemented")
}

// RemoveAttendee implements IAttendeeRepository.
func (a *attendeeRepository) RemoveAttendee(ctx context.Context, partyID string, userID string) error {
	panic("unimplemented")
}

// UpdateAttendeeStatus implements IAttendeeRepository.
func (a *attendeeRepository) UpdateAttendeeStatus(ctx context.Context, partyID string, userID string, status string) error {
	panic("unimplemented")
}

func NewAttendeeRepository(db *gorm.DB) IAttendeeRepository {
	return &attendeeRepository{db: db}
}

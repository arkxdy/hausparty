package repository

import (
	"context"
	models "hausparty/libs/db/models/users"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	CreateSession(ctx context.Context, session *models.Session) error
	GetSession(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

type sessionRepository struct {
	db *gorm.DB
}

// CreateSession implements ISessionRepository.
func (s *sessionRepository) CreateSession(ctx context.Context, session *models.Session) error {
	panic("unimplemented")
}

// DeleteSession implements ISessionRepository.
func (s *sessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	panic("unimplemented")
}

// GetSession implements ISessionRepository.
func (s *sessionRepository) GetSession(ctx context.Context, sessionID string) (*models.Session, error) {
	panic("unimplemented")
}

func NewSessionRepository(db *gorm.DB) ISessionRepository {
	return &sessionRepository{db: db}
}

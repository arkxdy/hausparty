package repository

import (
	"context"
	models "hausparty/libs/db/models/party"

	"gorm.io/gorm"
)

type IPartyRepository interface {
	CreateParty(ctx context.Context, party *models.Party) error
	GetPartyByID(ctx context.Context, id string) (*models.Party, error)
	FindNearbyParties(ctx context.Context, lat, lon float64, radius int) ([]models.Party, error)
	FindPaginatedNearbyParties(ctx context.Context, lat, lon float64, radius, page, pageSize int) ([]models.Party, error)
	UpdateParty(ctx context.Context, party *models.Party) error
	CancelParty(ctx context.Context, partyID string) error
}

type partyRepository struct {
	db *gorm.DB
}

// CancelParty implements IPartyRepository.
func (p *partyRepository) CancelParty(ctx context.Context, partyID string) error {
	panic("unimplemented")
}

// CreateParty implements IPartyRepository.
func (p *partyRepository) CreateParty(ctx context.Context, party *models.Party) error {
	panic("unimplemented")
}

// FindNearbyParties implements IPartyRepository.
func (p *partyRepository) FindNearbyParties(ctx context.Context, lat float64, lon float64, radius int) ([]models.Party, error) {
	var parties []models.Party
	err := p.db.WithContext(ctx).
		Where("ST_DWithin(location, ST_MakePoint(?, ?)::geography, ?)", lon, lat, radius).
		Find(&parties).
		Error
	return parties, err
}

func (p *partyRepository) FindPaginatedNearbyParties(ctx context.Context, lat, lon float64, radius, page, pageSize int) ([]models.Party, error) {
	var parties []models.Party
	err := p.db.WithContext(ctx).
		Where("ST_DWithin(location, ST_MakePoint(?, ?)::geography, ?)", lon, lat, radius).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&parties).
		Error
	return parties, err
}

// GetPartyByID implements IPartyRepository.
func (p *partyRepository) GetPartyByID(ctx context.Context, id string) (*models.Party, error) {
	var party models.Party
	err := p.db.WithContext(ctx).
		Preload("Attendees").
		Preload("Reports").
		First(&party, "id = ?", id).
		Error
	return &party, err
}

// UpdateParty implements IPartyRepository.
func (p *partyRepository) UpdateParty(ctx context.Context, party *models.Party) error {
	panic("unimplemented")
}

func NewPartyRepository(db *gorm.DB) IPartyRepository {
	return &partyRepository{db: db}
}

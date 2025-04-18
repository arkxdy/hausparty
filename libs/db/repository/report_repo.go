package repository

import (
	"context"
	models "hauparty/libs/db/models/party"

	"gorm.io/gorm"
)

type IReportRepository interface {
	CreateReport(ctx context.Context, report *models.Report) error
	GetReports(ctx context.Context, status string) ([]models.Report, error)
	ResolveReport(ctx context.Context, reportID string) error
}

type reportRepository struct {
	db *gorm.DB
}

// CreateReport implements IReportRepository.
func (r reportRepository) CreateReport(ctx context.Context, report *models.Report) error {
	panic("unimplemented")
}

// GetReports implements IReportRepository.
func (r reportRepository) GetReports(ctx context.Context, status string) ([]models.Report, error) {
	panic("unimplemented")
}

// ResolveReport implements IReportRepository.
func (r reportRepository) ResolveReport(ctx context.Context, reportID string) error {
	panic("unimplemented")
}

func NewReportRepository(db *gorm.DB) IReportRepository {
	return reportRepository{db: db}
}

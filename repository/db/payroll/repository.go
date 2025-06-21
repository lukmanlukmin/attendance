// Package payroll ...
package payroll

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	"context"

	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// IPayroll ...
type IPayroll interface {
	Create(ctx context.Context, p *model.Payroll) error
	GetByAttendacePeriod(ctx context.Context, attendancePeriodID uuid.UUID, status *string) ([]model.Payroll, error)
	Update(ctx context.Context, p *model.Payroll) error
}

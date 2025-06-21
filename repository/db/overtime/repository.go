// Package overtime ...
package overtime

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	model "attendance/model/db"
	"context"
	"time"

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

// IOvertime ...
type IOvertime interface {
	Create(ctx context.Context, ot *model.Overtime) error
	IsOvertimeSubmitted(ctx context.Context, employeeID uuid.UUID, date time.Time) (bool, error)
	SumEmployeeOvertimeHours(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error)
}

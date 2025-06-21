// Package attendanceperiod ...
package attendanceperiod

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

// IAttendancePeriod ...
type IAttendancePeriod interface {
	Create(ctx context.Context, ap *model.AttendancePeriod) error
	IsOverLapping(ctx context.Context, startDate, endDate time.Time) (bool, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*model.AttendancePeriod, error)
}

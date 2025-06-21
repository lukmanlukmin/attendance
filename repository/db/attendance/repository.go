// Package attendance ...
package attendance

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

// IAttendance ...
type IAttendance interface {
	Create(ctx context.Context, a *model.Attendance) error
	IsAttendanceSubmitted(ctx context.Context, employeeID uuid.UUID, date time.Time) (bool, error)
	CountEmployeeAttendance(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error)
}

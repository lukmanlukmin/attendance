// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// AttendancePeriod ...
type AttendancePeriod struct {
	ID        uuid.UUID  `db:"id"`
	StartDate time.Time  `db:"start_date"`
	EndDate   time.Time  `db:"end_date"`
	CreatedBy *uuid.UUID `db:"created_by"`
	CreatedIP *string    `db:"created_ip"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	IsDeleted bool       `db:"is_deleted"`
}

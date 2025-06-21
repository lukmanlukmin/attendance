// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Payroll ...
type Payroll struct {
	ID                 uuid.UUID  `db:"id"`
	AttendancePeriodID uuid.UUID  `db:"attendance_period_id"`
	Status             string     `db:"status"`
	CreatedBy          *uuid.UUID `db:"created_by"`
	CreatedIP          *string    `db:"created_ip"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
	IsDeleted          bool       `db:"is_deleted"`
}

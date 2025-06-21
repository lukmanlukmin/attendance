// Package payload ...
package payload

import "time"

// CreateAttendancePeriodRequest ...
type CreateAttendancePeriodRequest struct {
	StartDate time.Time `json:"start_date" validate:"required" example:"2025-07-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" validate:"required" example:"2025-07-31T23:59:59Z"`
}

// Package event ...
package event

import "github.com/google/uuid"

// CalculatePayrollJob ...
type CalculatePayrollJob struct {
	AttendancePeriodID uuid.UUID `json:"attendance_period_id"`
	PayrollID          uuid.UUID `json:"payroll_id"`
}

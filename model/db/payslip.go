// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Payslip ...
type Payslip struct {
	ID                   uuid.UUID  `db:"id"`
	PayrollID            uuid.UUID  `db:"payroll_id"`
	EmployeeID           uuid.UUID  `db:"employee_id"`
	BaseSalary           int        `db:"base_salary"`
	AttendanceDays       int        `db:"attendance_days"`
	AttendanceOffDays    int        `db:"attendance_off_days"`
	AttendanceDeduction  int        `db:"attendance_deduction"`
	OvertimeHours        int        `db:"overtime_hours"`
	OvertimeMultiplyRate float64    `db:"overtime_multiply_rate"`
	OvertimePay          int        `db:"overtime_pay"`
	ReimbursementTotal   int        `db:"reimbursement_total"`
	TakeHomePay          int        `db:"take_home_pay"`
	GeneratedAt          time.Time  `db:"generated_at"`
	DeletedAt            *time.Time `db:"deleted_at"`
	IsDeleted            bool       `db:"is_deleted"`
}

// ResumePayslip ...
type ResumePayslip struct {
	ID               uuid.UUID `db:"id"`
	EmployeeID       uuid.UUID `db:"employee_id"`
	EmployeeFullName string    `db:"full_name"`
	TakeHomePay      int       `db:"take_home_pay"`
}

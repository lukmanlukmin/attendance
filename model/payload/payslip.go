// Package payload ...
package payload

import "github.com/google/uuid"

// PayslipDetailResponse ...
type PayslipDetailResponse struct {
	ID                   uuid.UUID                `json:"id"`
	EmployeeID           uuid.UUID                `json:"employee_id"`
	BaseSalary           int                      `json:"base_salary"`
	AttendanceDays       int                      `json:"attendance_days"`
	AttendanceOffDays    int                      `json:"attendance_off_days"`
	AttendanceDeduction  int                      `json:"attendance_deduction"`
	OvertimeHours        int                      `json:"overtime_hours"`
	OvertimeMultiplyRate float64                  `json:"overtime_multiply_rate"`
	OvertimePay          int                      `json:"overtime_pay"`
	Reimbursements       []ReimbursementBreakdown `json:"reimbursements"`
	ReimbursementTotal   int                      `json:"reimbursement_total"`
	TakeHomePay          int                      `json:"take_home_pay"`
}

// ReimbursementBreakdown ...
type ReimbursementBreakdown struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

// PayslipSummaryItem ...
type PayslipSummaryItem struct {
	EmployeeID   uuid.UUID `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	TakeHomePay  int       `json:"take_home_pay"`
}

// PayslipSummaryResponse ...
type PayslipSummaryResponse struct {
	Page          uint64               `json:"page"`
	PerPage       uint64               `json:"per_page"`
	TotalData     int                  `json:"total_data"`
	TotalTakeHome int                  `json:"total_take_home"`
	Data          []PayslipSummaryItem `json:"data"`
}

package constant

import (
	"errors"
	"net/http"
)

var (
	// ErrInvalidUserCredentials ...
	ErrInvalidUserCredentials = errors.New("invalid user credentials")
	// ErrFailedGenerateToken ...
	ErrFailedGenerateToken = errors.New("failed to generate token")
	// ErrFailedValidateTokenID ...
	ErrFailedValidateTokenID = errors.New("failed to validate token id")

	// ErrOverlappingAttendancePeriod ...
	ErrOverlappingAttendancePeriod = errors.New("overlapping attendance period")
	// ErrWeekendAttendance ...
	ErrWeekendAttendance = errors.New("cannot submit attendance on weekends")
	// ErrEmployeeNotFound ...
	ErrEmployeeNotFound = errors.New("employee not found")
	// ErrAlreadySubmitAttendance ...
	ErrAlreadySubmitAttendance = errors.New("already submitted attendance")
	// ErrMaximumOvertime ...
	ErrMaximumOvertime = errors.New("overtime cannot exceed maximum hours per day")
	// ErrSubmitOvertimeBeforeWorkingHour ...
	ErrSubmitOvertimeBeforeWorkingHour = errors.New("can't submit overtime after working hours")
	// ErrAlreadySubmitOvertime ...
	ErrAlreadySubmitOvertime = errors.New("overtime already submitted for today")
	// ErrPayrollNotFound ...
	ErrPayrollNotFound = errors.New("payroll not found")
	// ErrPayrollAlreadyProcessed ...
	ErrPayrollAlreadyProcessed = errors.New("payroll already processed")
	// ErrPayrollAlreadyRequested ...
	ErrPayrollAlreadyRequested = errors.New("payroll already requested")
	// ErrAttendancePeriodNotFound ...
	ErrAttendancePeriodNotFound = errors.New("attendance period not found")
	// ErrFailedCountEmployeeAttendance ...
	ErrFailedCountEmployeeAttendance = errors.New("failed to count employee attendance")
	// ErrFailedSumEmployeeOvertimeHours ...
	ErrFailedSumEmployeeOvertimeHours = errors.New("failed to sum employee overtime hours")
	// ErrFailedSumEmployeeReimbursement ...
	ErrFailedSumEmployeeReimbursement = errors.New("failed to sum employee reimbursement")
	// ErrPayslipNotFound ...
	ErrPayslipNotFound = errors.New("payslip not found")
	// ErrFailedGetReimbursement ...
	ErrFailedGetReimbursement = errors.New("failed to get reimbursement")
)

var errorStatusMap = map[error]int{
	ErrInvalidUserCredentials:          http.StatusBadRequest,
	ErrFailedGenerateToken:             http.StatusBadRequest,
	ErrFailedValidateTokenID:           http.StatusBadRequest,
	ErrOverlappingAttendancePeriod:     http.StatusUnprocessableEntity,
	ErrWeekendAttendance:               http.StatusUnprocessableEntity,
	ErrEmployeeNotFound:                http.StatusBadRequest,
	ErrAlreadySubmitAttendance:         http.StatusUnprocessableEntity,
	ErrMaximumOvertime:                 http.StatusBadRequest,
	ErrSubmitOvertimeBeforeWorkingHour: http.StatusBadRequest,
	ErrAlreadySubmitOvertime:           http.StatusBadRequest,
	ErrPayrollNotFound:                 http.StatusBadRequest,
	ErrPayrollAlreadyProcessed:         http.StatusBadRequest,
	ErrPayrollAlreadyRequested:         http.StatusBadRequest,
	ErrAttendancePeriodNotFound:        http.StatusBadRequest,
	ErrFailedCountEmployeeAttendance:   http.StatusBadRequest,
	ErrFailedSumEmployeeOvertimeHours:  http.StatusBadRequest,
	ErrFailedSumEmployeeReimbursement:  http.StatusBadRequest,
	ErrPayslipNotFound:                 http.StatusBadRequest,
	ErrFailedGetReimbursement:          http.StatusBadRequest,
}

// GetHTTPStatus ...
func GetHTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	for e, status := range errorStatusMap {
		if errors.Is(err, e) {
			return status
		}
	}

	return http.StatusInternalServerError
}

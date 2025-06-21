package payroll

import (
	"attendance/constant"
	"attendance/model/payload"
	"attendance/utils"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// GetPayslip ...
func (s *Service) GetPayslip(ctx context.Context, payrollID uuid.UUID) (*payload.PayslipDetailResponse, error) {
	userCtx := utils.GetUserContext(ctx)
	if !userCtx.EmployeeID.Valid {
		return nil, constant.ErrEmployeeNotFound
	}
	payslip, err := s.Repository.DB.Payslip.GetByPayrollAndEmployee(ctx, payrollID, userCtx.EmployeeID.UUID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if payslip == nil {
		return nil, constant.ErrPayslipNotFound
	}

	reimbursement, err := s.Repository.DB.Reimbursement.GetByEmployeeAndPeriod(ctx, userCtx.EmployeeID.UUID, payslip.PayrollID)
	if err != nil {
		return nil, constant.ErrFailedGetReimbursement
	}
	resultReimbursement := []payload.ReimbursementBreakdown{}
	for _, r := range reimbursement {
		resultReimbursement = append(resultReimbursement, payload.ReimbursementBreakdown{
			Amount:      r.Amount,
			Description: r.Description,
		})
	}

	result := &payload.PayslipDetailResponse{
		ID:                   payslip.ID,
		EmployeeID:           payslip.EmployeeID,
		BaseSalary:           payslip.BaseSalary,
		AttendanceDays:       payslip.AttendanceDays,
		AttendanceOffDays:    payslip.AttendanceOffDays,
		AttendanceDeduction:  payslip.AttendanceDeduction,
		OvertimeHours:        payslip.OvertimeHours,
		OvertimeMultiplyRate: payslip.OvertimeMultiplyRate,
		OvertimePay:          payslip.OvertimePay,
		ReimbursementTotal:   payslip.ReimbursementTotal,
		TakeHomePay:          payslip.TakeHomePay,
		Reimbursements:       resultReimbursement,
	}
	return result, nil
}

package payroll

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/utils"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

// GeneratePayslip ...
func (s *Service) GeneratePayslip(ctx context.Context, period *db.AttendancePeriod, employeeID, payrollID uuid.UUID) (db.Payslip, error) {

	payslip := db.Payslip{
		PayrollID:  payrollID,
		EmployeeID: employeeID,
	}

	employee, err := s.Repository.DB.Employee.GetByID(ctx, employeeID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return payslip, err
	}
	if employee == nil {
		return payslip, constant.ErrEmployeeNotFound
	}
	payslip.BaseSalary = employee.Salary

	// attendance
	attCount, err := s.Repository.DB.Attendance.CountEmployeeAttendance(ctx, employeeID, period.ID)
	if err != nil {
		return payslip, constant.ErrFailedCountEmployeeAttendance
	}
	payslip.AttendanceDays = attCount

	// overtime
	overtimeHours, err := s.Repository.DB.Overtime.SumEmployeeOvertimeHours(ctx, employeeID, period.ID)
	if err != nil {
		return payslip, constant.ErrFailedSumEmployeeOvertimeHours
	}
	payslip.OvertimeHours = overtimeHours

	// reimbursement
	reimbTotal, err := s.Repository.DB.Reimbursement.SumEmployeeReimbursements(ctx, employeeID, period.ID)
	if err != nil {
		return payslip, constant.ErrFailedSumEmployeeReimbursement
	}
	payslip.ReimbursementTotal = reimbTotal

	// calculate
	totalWorkingDays := utils.CountWeekdays(period.StartDate, period.EndDate)
	dailySalary := employee.Salary / totalWorkingDays
	if attCount < totalWorkingDays {
		payslip.AttendanceOffDays = totalWorkingDays - attCount
		payslip.AttendanceDeduction = payslip.AttendanceOffDays * dailySalary
	}

	overtimePay := utils.CalculateOvertimePay(employee.Salary, overtimeHours, s.Config.Application.MultiplyOvertimeRate)
	payslip.OvertimePay = overtimePay
	payslip.TakeHomePay = payslip.AttendanceDays - payslip.AttendanceDeduction + payslip.OvertimePay + payslip.ReimbursementTotal

	return payslip, nil
}

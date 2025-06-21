package payslip

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// CreateBulk ...
func (r *Repository) CreateBulk(ctx context.Context, payslips []model.Payslip) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	if len(payslips) == 0 {
		return nil
	}

	builder := sq.Insert("payslips").
		Columns(
			"id", "payroll_id", "employee_id",
			"base_salary", "attendance_days", "attendance_off_days", "attendance_deduction",
			"overtime_hours", "overtime_multiply_rate", "overtime_pay",
			"reimbursement_total", "take_home_pay", "generated_at",
		)

	for _, p := range payslips {
		p.ID = uuid.New()
		builder = builder.Values(
			p.ID, p.PayrollID, p.EmployeeID,
			p.BaseSalary, p.AttendanceDays, p.AttendanceOffDays, p.AttendanceDeduction,
			p.OvertimeHours, p.OvertimeMultiplyRate, p.OvertimePay,
			p.ReimbursementTotal, p.TakeHomePay, sq.Expr("NOW()"),
		)
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	if _, err = db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

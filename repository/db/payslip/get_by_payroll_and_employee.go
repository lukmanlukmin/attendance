package payslip

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// GetByPayrollAndEmployee ...
func (r *Repository) GetByPayrollAndEmployee(ctx context.Context, payrollID, employeeID uuid.UUID) (*model.Payslip, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := &model.Payslip{}
	query, args, err := sq.
		Select("p.id", "p.payroll_id", "p.employee_id", "p.base_salary", "p.attendance_days", "p.attendance_off_days", "p.attendance_deduction", "p.overtime_hours", "p.overtime_multiply_rate", "p.overtime_pay", "p.reimbursement_total", "p.take_home_pay", "p.generated_at", "p.deleted_at", "p.is_deleted").
		From("payslips p").
		Join("payrolls pr ON pr.id = p.payroll_id").
		Where(sq.Eq{
			"p.employee_id": employeeID,
			"p.is_deleted":  false,
			"pr.id":         payrollID,
		}).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err = db.GetContext(ctx, data, query, args...); err != nil {
		return nil, err
	}
	return data, nil

}

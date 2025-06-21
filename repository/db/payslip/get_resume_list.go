// Package payslip ...
package payslip

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// GetResumeList ...
func (r *Repository) GetResumeList(ctx context.Context, payrollID uuid.UUID, page, perPage uint64) ([]model.ResumePayslip, int, error) {
	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	payslips := []model.ResumePayslip{}

	countQuery, countArgs, err := sq.
		Select("COUNT(*)").
		From("payslips p").
		Join("payrolls pr ON pr.id = p.payroll_id").
		Join("employees e ON e.id = p.employee_id").
		Where(sq.Eq{
			"p.is_deleted":  false,
			"pr.is_deleted": false,
			"pr.id":         payrollID,
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return payslips, 0, err
	}
	var total int
	if err := db.GetContext(ctx, &total, countQuery, countArgs...); err != nil {
		return payslips, 0, err
	}

	offset := (page - 1) * perPage
	query, args, err := sq.
		Select("p.id", "p.employee_id", "e.full_name", "p.take_home_pay").
		From("payslips p").
		Join("payrolls pr ON pr.id = p.payroll_id").
		Join("employees e ON e.id = p.employee_id").
		Where(sq.Eq{
			"p.is_deleted":  false,
			"pr.is_deleted": false,
			"pr.id":         payrollID,
		}).
		Limit(perPage).
		Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return payslips, 0, err
	}

	if err := db.SelectContext(ctx, &payslips, query, args...); err != nil {
		return payslips, 0, err
	}
	return payslips, total, nil
}

package payslip

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// GetTotalTakeHomePay ...
func (r *Repository) GetTotalTakeHomePay(ctx context.Context, payrollID uuid.UUID) (int, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("COALESCE(SUM(p.take_home_pay), 0)").
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
		return 0, err
	}

	var total int
	if err := db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, err
	}

	return total, nil

}

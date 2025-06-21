// Package reimbursement ...
package reimbursement

import (
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// SumEmployeeReimbursements ...
func (r *Repository) SumEmployeeReimbursements(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("COALESCE(SUM(r.amount), 0)").
		From("reimbursements r").
		Join("attendance_periods ap ON r.date BETWEEN ap.start_date AND ap.end_date").
		Where(sq.Eq{"r.employee_id": employeeID}).
		Where(sq.Eq{"ap.id": attendancePeriodID}).
		Where(sq.Eq{"r.is_deleted": false}).
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

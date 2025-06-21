// Package overtime ...
package overtime

import (
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// SumEmployeeOvertimeHours ...
func (r *Repository) SumEmployeeOvertimeHours(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("COALESCE(SUM(o.hours), 0)").
		From("overtimes o").
		Join("attendance_periods ap ON o.date BETWEEN ap.start_date AND ap.end_date").
		Where(sq.Eq{
			"o.employee_id": employeeID,
			"ap.id":         attendancePeriodID,
			"o.is_deleted":  false,
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

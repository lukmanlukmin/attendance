// Package attendance ...
package attendance

import (
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// CountEmployeeAttendance ...
func (r *Repository) CountEmployeeAttendance(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("COUNT(1)").
		From("attendances a").
		Join("attendance_periods ap ON a.date BETWEEN ap.start_date AND ap.end_date").
		Where(sq.Eq{
			"a.employee_id": employeeID,
			"ap.id":         attendancePeriodID,
			"a.is_deleted":  false,
		}).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return 0, err
	}

	var count int
	if err := db.GetContext(ctx, &count, query, args...); err != nil {
		return 0, err
	}

	return count, nil
}

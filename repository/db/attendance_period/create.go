// Package attendanceperiod ...
package attendanceperiod

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// Create ...
func (r *Repository) Create(ctx context.Context, ap *model.AttendancePeriod) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	ap.ID = uuid.New()
	query, args, err := sq.
		Insert("attendance_periods").
		SetMap(sq.Eq{
			"id":         ap.ID,
			"start_date": ap.StartDate,
			"end_date":   ap.EndDate,
			"created_by": ap.CreatedBy,
			"created_ip": ap.CreatedIP,
			"created_at": sq.Expr("NOW()"),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	if _, err = db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

// Package attendance ...
package attendance

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// Create ...
func (r *Repository) Create(ctx context.Context, a *model.Attendance) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	a.ID = uuid.New()
	query, args, err := sq.
		Insert("attendances").
		SetMap(sq.Eq{
			"id":          a.ID,
			"employee_id": a.EmployeeID,
			"date":        a.Date,
			"created_by":  a.CreatedBy,
			"created_ip":  a.CreatedIP,
			"created_at":  sq.Expr("NOW()"),
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

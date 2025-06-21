// Package overtime ...
package overtime

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// Create ...
func (r *Repository) Create(ctx context.Context, ot *model.Overtime) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	ot.ID = uuid.New()
	query, args, err := sq.
		Insert("overtimes").
		SetMap(sq.Eq{
			"id":          ot.ID,
			"employee_id": ot.EmployeeID,
			"date":        ot.Date,
			"hours":       ot.Hours,
			"created_by":  ot.CreatedBy,
			"created_ip":  ot.CreatedIP,
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

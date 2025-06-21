// Package payroll ...
package payroll

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// Create ...
func (r *Repository) Create(ctx context.Context, p *model.Payroll) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	p.ID = uuid.New()
	query, args, err := sq.
		Insert("payrolls").
		SetMap(sq.Eq{
			"id":                   p.ID,
			"attendance_period_id": p.AttendancePeriodID,
			"status":               p.Status,
			"created_by":           p.CreatedBy,
			"created_ip":           p.CreatedIP,
			"created_at":           sq.Expr("NOW()"),
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

// Package payroll ...
package payroll

import (
	models "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/lukmanlukmin/go-lib/database"
)

// Update ...
func (r *Repository) Update(ctx context.Context, p *models.Payroll) error {
	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Update("payrolls").
		SetMap(sq.Eq{
			"status":     p.Status,
			"updated_at": sq.Expr("NOW()"),
			"deleted_at": p.DeletedAt,
			"is_deleted": p.IsDeleted,
		}).
		Where(sq.Eq{"id": p.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	return err
}

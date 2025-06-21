// Package role ...
package role

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// Create ...
func (r *Repository) Create(ctx context.Context, role *model.Role) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	role.ID = uuid.New()
	query, args, err := sq.
		Insert("roles").
		SetMap(sq.Eq{
			"id":         role.ID,
			"name":       role.Name,
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

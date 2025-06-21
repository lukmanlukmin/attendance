// Package employee ...
package employee

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// Create ...
func (r *Repository) Create(ctx context.Context, emp *model.Employee) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	emp.ID = uuid.New()
	query, args, err := sq.
		Insert("employees").
		SetMap(sq.Eq{
			"id":         emp.ID,
			"user_id":    emp.UserID,
			"full_name":  emp.FullName,
			"salary":     emp.Salary,
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

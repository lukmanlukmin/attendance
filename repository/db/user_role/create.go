// Package userrole ...
package userrole

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// Create ...
func (r *Repository) Create(ctx context.Context, userRole *model.UserRole) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	userRole.ID = uuid.New()
	query, args, err := sq.
		Insert("user_roles").
		SetMap(sq.Eq{
			"id":         userRole.ID,
			"user_id":    userRole.UserID,
			"role_id":    userRole.RoleID,
			"created_by": userRole.CreatedBy,
			"created_ip": userRole.CreatedIP,
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

// Package role ...
package role

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/lukmanlukmin/go-lib/database"
)

// GetByUserID ...
func (r *Repository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Role, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := []model.Role{}
	query, args, err := sq.
		Select("r.id", "r.name", "r.created_at", "r.updated_at").
		From("roles r").
		Join("user_roles ur ON ur.role_id = r.id").
		Where(sq.Eq{
			"ur.user_id":    userID,
			"ur.is_deleted": false,
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return data, err
	}
	if err = db.SelectContext(ctx, &data, query, args...); err != nil {
		return data, err
	}
	return data, nil
}

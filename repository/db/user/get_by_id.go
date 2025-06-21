// Package user ...
package user

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// GetByID ...
func (r *Repository) GetByID(ctx context.Context, ID uuid.UUID) (*model.User, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := &model.User{}
	query, args, err := sq.
		Select("id", "username", "password", "created_by", "updated_by", "created_ip", "updated_ip", "created_at", "updated_at", "deleted_at", "is_deleted").
		From("users").
		Where(sq.Eq{
			"id":         ID,
			"is_deleted": false,
		}).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err = db.GetContext(ctx, data, query, args...); err != nil {
		return nil, err
	}
	return data, nil
}

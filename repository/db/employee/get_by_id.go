// Package employee ...
package employee

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// GetByID ...
func (r *Repository) GetByID(ctx context.Context, ID uuid.UUID) (*model.Employee, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := &model.Employee{}
	query, args, err := sq.
		Select("id", "user_id", "full_name", "salary", "created_at", "updated_at", "deleted_at", "is_deleted").
		From("employees").
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

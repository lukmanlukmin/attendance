// Package role ...
package role

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/lukmanlukmin/go-lib/database"
)

// GetByName ...
func (r *Repository) GetByName(ctx context.Context, name string) (*model.Role, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := &model.Role{}
	query, args, err := sq.
		Select("id", "name", "created_at", "updated_at").
		From("roles").
		Where("name = ?", name).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	if err = db.GetContext(ctx, data, query, args...); err != nil {
		return nil, err
	}
	return data, nil
}

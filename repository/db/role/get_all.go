// Package role ...
package role

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/lukmanlukmin/go-lib/database"
)

// GetAll ...
func (r *Repository) GetAll(ctx context.Context) ([]model.Role, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := []model.Role{}
	query, args, err := sq.
		Select("id", "name", "created_at", "updated_at").
		From("roles").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return data, err
	}
	if err = db.SelectContext(ctx, &data, query, args...); err != nil {
		return data, err
	}
	return data, nil
}

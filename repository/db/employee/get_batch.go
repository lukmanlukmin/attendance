// Package employee ...
package employee

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/lukmanlukmin/go-lib/database"
)

// GetBatch ...
func (r *Repository) GetBatch(ctx context.Context, batchSize, offset int) ([]model.Employee, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := []model.Employee{}
	query, args, err := sq.
		Select("id", "user_id", "full_name", "salary", "created_at", "updated_at", "deleted_at", "is_deleted").
		From("roles").
		Where(sq.Eq{"is_deleted": false}).
		OrderBy("id").
		Limit(uint64(batchSize)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return data, err
	}
	if err = db.SelectContext(ctx, &data, query, args...); err != nil {
		return data, err
	}
	return data, nil
}

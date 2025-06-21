// Package reimbursement ...
package reimbursement

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// Create ...
func (r *Repository) Create(ctx context.Context, ri *model.Reimbursement) error {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	ri.ID = uuid.New()
	query, args, err := sq.
		Insert("roles").
		SetMap(sq.Eq{
			"id":          ri.ID,
			"employee_id": ri.EmployeeID,
			"date":        ri.Date,
			"amount":      ri.Amount,
			"description": ri.Description,
			"created_by":  ri.CreatedBy,
			"created_ip":  ri.CreatedIP,
			"created_at":  sq.Expr("NOW()"),
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

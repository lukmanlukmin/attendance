// Package reimbursement ...
package reimbursement

import (
	model "attendance/model/db"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
)

// GetByEmployeeAndPeriod ...
func (r *Repository) GetByEmployeeAndPeriod(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) ([]model.Reimbursement, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	data := []model.Reimbursement{}
	query, args, err := sq.Select("r.id", "r.employee_id", "r.amount", "r.description", "r.created_at").
		From("reimbursements r").
		Join("attendance_periods ap ON ap.id = ?", attendancePeriodID).
		Where(sq.Eq{
			"r.employee_id": employeeID,
			"r.is_deleted":  false,
			"ap.is_deleted": false,
		}).
		Where("r.date BETWEEN ap.start_date AND ap.end_date").
		OrderBy("r.created_at ASC").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err = db.SelectContext(ctx, &data, query, args...); err != nil {
		return nil, err
	}
	return data, nil

}

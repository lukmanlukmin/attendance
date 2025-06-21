// Package payroll ...
package payroll

import (
	model "attendance/model/db"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// GetByAttendacePeriod ...
func (r *Repository) GetByAttendacePeriod(ctx context.Context, attendancePeriodID uuid.UUID, status *string) ([]model.Payroll, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}
	where := sq.Eq{
		"attendance_period_id": attendancePeriodID,
		"is_deleted":           false,
	}
	if status != nil {
		where["status"] = *status
	}

	data := []model.Payroll{}
	query, args, err := sq.
		Select("id", "attendance_period_id", "status", "created_by", "created_ip", "created_at", "updated_at", "deleted_at", "is_deleted").
		From("payrolls").
		Where(where).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err = db.SelectContext(ctx, &data, query, args...); err != nil {
		return data, err
	}
	return data, nil
}

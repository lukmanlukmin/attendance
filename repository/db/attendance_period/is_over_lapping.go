// Package attendanceperiod ...
package attendanceperiod

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// IsOverLapping ...
func (r *Repository) IsOverLapping(ctx context.Context, startDate, endDate time.Time) (bool, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("1").
		From("attendance_periods").
		Where("is_deleted = FALSE").
		Where("NOT (? < start_date OR ? > end_date)", endDate, startDate).
		Limit(1).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return false, err
	}

	var found int
	err = db.GetContext(ctx, &found, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return found == 1, nil
}

// Package attendance ...
package attendance

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"

	"context"

	sq "github.com/Masterminds/squirrel"
)

// IsAttendanceSubmitted ...
func (r *Repository) IsAttendanceSubmitted(ctx context.Context, employeeID uuid.UUID, date time.Time) (bool, error) {

	var db database.SQLQueryExec = r.DB
	if tx := database.GetTxFromContext(ctx); tx != nil {
		db = tx
	}

	query, args, err := sq.
		Select("1").
		From("attendances").
		Where(sq.Eq{
			"employee_id": employeeID,
			"date":        date.Format("2006-01-02"),
			"is_deleted":  false,
		}).
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

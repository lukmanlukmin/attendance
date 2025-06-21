// Package attendance ...
package attendance

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestIsAttendanceSubmitted ...
func TestIsAttendanceSubmitted(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockEid := uuid.New()
	mockDate := time.Now()

	tests := []struct {
		name     string
		eid      uuid.UUID
		d        time.Time
		mockFunc func(e uuid.UUID, d time.Time)
		wantErr  bool
	}{
		{
			name: "success",
			eid:  mockEid,
			d:    mockDate,
			mockFunc: func(e uuid.UUID, p time.Time) {
				rows := sqlmock.NewRows([]string{
					"1",
				}).AddRow(
					1,
				)

				mocks.ExpectQuery(`SELECT 1 FROM attendances WHERE date = $1 AND employee_id = $2 AND is_deleted = $3 LIMIT 1`).
					WithArgs(p.Format("2006-01-02"), e, false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			eid:  mockEid,
			d:    mockDate,
			mockFunc: func(e uuid.UUID, p time.Time) {
				mocks.ExpectQuery(`SELECT 1 FROM attendances WHERE date = $1 AND employee_id = $2 AND is_deleted = $3 LIMIT 1`).
					WithArgs(p.Format("2006-01-02"), e, false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.eid, tt.d)
			result, err := r.IsAttendanceSubmitted(context.Background(), tt.eid, tt.d)
			if tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, false, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, true, result)
			}
		})
	}
}

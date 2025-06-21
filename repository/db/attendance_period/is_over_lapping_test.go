// Package attendanceperiod ...
package attendanceperiod

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestIsOverLapping ...
func TestIsOverLapping(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	// mockEid := uuid.New()
	mockDateStart := time.Now()
	mockDateEnd := time.Now().Add(time.Hour * 24 * 25)

	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		mockFunc func(s, e time.Time)
		wantErr  bool
	}{
		{
			name:  "success",
			start: mockDateStart,
			end:   mockDateEnd,
			mockFunc: func(s, e time.Time) {
				rows := sqlmock.NewRows([]string{
					"1",
				}).AddRow(
					1,
				)

				mocks.ExpectQuery(`SELECT 1 FROM attendance_periods WHERE is_deleted = FALSE AND NOT ($1 < start_date OR $2 > end_date) LIMIT 1`).
					WithArgs(e, s).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "error",
			start: mockDateStart,
			end:   mockDateEnd,
			mockFunc: func(s, e time.Time) {
				mocks.ExpectQuery(`SELECT 1 FROM attendance_periods WHERE is_deleted = FALSE AND NOT ($1 < start_date OR $2 > end_date) LIMIT 1`).
					WithArgs(e, s).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.start, tt.end)
			result, err := r.IsOverLapping(context.Background(), tt.start, tt.end)
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

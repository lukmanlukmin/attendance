// Package payroll ...
package payroll

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

// TestGetByAttendacePeriod ...
func TestGetByAttendacePeriod(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockStatus := "pending"
	tests := []struct {
		name     string
		ap       uuid.UUID
		st       *string
		mockFunc func(ap uuid.UUID, st *string)
		wantErr  bool
	}{
		{
			name: "success",
			ap:   uuid.New(),
			st:   &mockStatus,
			mockFunc: func(ap uuid.UUID, st *string) {
				rows := sqlmock.NewRows([]string{
					"id", "attendance_period_id", "status", "created_by", "created_ip", "created_at", "updated_at", "deleted_at", "is_deleted",
				}).AddRow(
					uuid.New(), uuid.New(), mockStatus, nil, nil, time.Now(), time.Now(), nil, false,
				)

				mocks.ExpectQuery(`SELECT id, attendance_period_id, status, created_by, created_ip, created_at, updated_at, deleted_at, is_deleted FROM payrolls WHERE attendance_period_id = $1 AND is_deleted = $2 AND status = $3`).
					WithArgs(ap, false, st).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			ap:   uuid.New(),
			st:   &mockStatus,
			mockFunc: func(ap uuid.UUID, st *string) {
				mocks.ExpectQuery(`SELECT id, attendance_period_id, status, created_by, created_ip, created_at, updated_at, deleted_at, is_deleted FROM payrolls WHERE attendance_period_id = $1 AND is_deleted = $2 AND status = $3`).
					WithArgs(ap, false, st).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.ap, tt.st)
			result, err := r.GetByAttendacePeriod(context.Background(), tt.ap, tt.st)
			if tt.wantErr {
				assert.Error(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, 0, len(result))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

// Package attendanceperiod ...
package attendanceperiod

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

// TestGetByID ...
func TestGetByID(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockID := uuid.New()

	tests := []struct {
		name     string
		input    uuid.UUID
		mockFunc func(id uuid.UUID)
		wantErr  bool
	}{
		{
			name:  "success",
			input: mockID,
			mockFunc: func(id uuid.UUID) {
				rows := sqlmock.NewRows([]string{
					"id", "start_date", "end_date", "created_by", "created_ip", "created_at", "updated_at", "deleted_at", "is_deleted",
				}).AddRow(
					mockID, time.Now(), time.Now(), nil, nil, time.Now(), time.Now(), nil, false,
				)

				mocks.ExpectQuery(`SELECT id, start_date, end_date, created_by, created_ip, created_at, updated_at, deleted_at, is_deleted FROM attendance_periods WHERE id = $1 AND is_deleted = $2`).
					WithArgs(id, false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "not found",
			input: mockID,
			mockFunc: func(id uuid.UUID) {
				mocks.ExpectQuery(`SELECT id, username, password, created_by, updated_by, created_ip, updated_ip, created_at, updated_at, deleted_at, is_deleted FROM users WHERE id = $1 AND is_deleted = $2`).
					WithArgs(id, false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.input)
			result, err := r.GetByID(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.input, result.ID)
			}
		})
	}
}

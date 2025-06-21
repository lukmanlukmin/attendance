// Package employee ...
package employee

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

// TestGetAll ...
func TestGetAll(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")

	tests := []struct {
		name     string
		bSize    int
		offset   int
		mockFunc func()
		wantErr  bool
	}{
		{
			name:   "success",
			bSize:  10,
			offset: 0,
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{
					"id", "user_id", "full_name", "salary", "created_at", "updated_at", "deleted_at", "is_deleted",
				}).AddRow(
					uuid.New(), uuid.New(), "name", 10000, time.Now(), time.Now(), nil, false,
				)

				mocks.ExpectQuery(`SELECT id, user_id, full_name, salary, created_at, updated_at, deleted_at, is_deleted FROM roles WHERE is_deleted = $1 ORDER BY id LIMIT 10 OFFSET 0`).
					WithArgs(false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:   "not found",
			bSize:  10,
			offset: 0,
			mockFunc: func() {
				mocks.ExpectQuery(`SELECT id, user_id, full_name, salary, created_at, updated_at, deleted_at, is_deleted FROM roles WHERE is_deleted = $1 ORDER BY id LIMIT 10 OFFSET 0`).
					WithArgs(false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc()
			result, err := r.GetBatch(context.Background(), tt.bSize, tt.offset)
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

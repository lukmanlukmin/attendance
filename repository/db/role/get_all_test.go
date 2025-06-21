// Package role ...
package role

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
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{
					"id", "name", "created_at", "updated_at",
				}).AddRow(
					uuid.New(), "name", time.Now(), time.Now(),
				)

				mocks.ExpectQuery(`SELECT id, name, created_at, updated_at FROM roles`).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			mockFunc: func() {
				mocks.ExpectQuery(`SELECT id, name, created_at, updated_at FROM roles`).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc()
			result, err := r.GetAll(context.Background())
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

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

// TestGetByUserID ...
func TestGetByUserID(t *testing.T) {
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
		mockFunc func(un uuid.UUID)
		wantErr  bool
	}{
		{
			name:  "success",
			input: mockID,
			mockFunc: func(un uuid.UUID) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "created_at", "updated_at",
				}).AddRow(
					uuid.New(), "name", time.Now(), time.Now(),
				)

				mocks.ExpectQuery(`SELECT r.id, r.name, r.created_at, r.updated_at FROM roles r JOIN user_roles ur ON ur.role_id = r.id WHERE ur.is_deleted = $1 AND ur.user_id = $2`).
					WithArgs(false, un).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "not found",
			input: mockID,
			mockFunc: func(un uuid.UUID) {
				mocks.ExpectQuery(`SELECT r.id, r.name, r.created_at, r.updated_at FROM roles r JOIN user_roles ur ON ur.role_id = r.id WHERE ur.is_deleted = $1 AND ur.user_id = $2`).
					WithArgs(false, un).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.input)
			result, err := r.GetByUserID(context.Background(), tt.input)
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

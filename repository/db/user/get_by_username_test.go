// Package user ...
package user

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
func TestGetByUsername(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockUname := "username"

	tests := []struct {
		name     string
		input    string
		mockFunc func(un string)
		wantErr  bool
	}{
		{
			name:  "success",
			input: mockUname,
			mockFunc: func(un string) {
				rows := sqlmock.NewRows([]string{
					"id", "username", "password", "created_by", "updated_by", "created_ip", "updated_ip", "created_at", "updated_at", "deleted_at", "is_deleted",
				}).AddRow(
					uuid.New(), "username", "password", nil, nil, nil, nil, time.Now(), time.Now(), nil, false,
				)

				mocks.ExpectQuery(`SELECT id, username, password, created_by, updated_by, created_ip, updated_ip, created_at, updated_at, deleted_at, is_deleted FROM users WHERE is_deleted = $1 AND username = $2`).
					WithArgs(false, un).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "not found",
			input: mockUname,
			mockFunc: func(un string) {
				mocks.ExpectQuery(`SELECT id, username, password, created_by, updated_by, created_ip, updated_ip, created_at, updated_at, deleted_at, is_deleted FROM users WHERE is_deleted = $1 AND username = $2`).
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
			result, err := r.GetByUsername(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.input, result.Username)
			}
		})
	}
}

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

// TestGetByName ...
func TestGetByName(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockName := "name"

	tests := []struct {
		name     string
		input    string
		mockFunc func(un string)
		wantErr  bool
	}{
		{
			name:  "success",
			input: mockName,
			mockFunc: func(un string) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "created_at", "updated_at",
				}).AddRow(
					uuid.New(), mockName, time.Now(), time.Now(),
				)

				mocks.ExpectQuery(`SELECT id, name, created_at, updated_at FROM roles WHERE name = $1`).
					WithArgs(un).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:  "not found",
			input: mockName,
			mockFunc: func(un string) {
				mocks.ExpectQuery(`SELECT id, name, created_at, updated_at FROM roles WHERE name = $1`).
					WithArgs(un).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.input)
			result, err := r.GetByName(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.input, result.Name)
			}
		})
	}
}

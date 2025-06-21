// Package userrole ...
package userrole

import (
	models "attendance/model/db"
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestCreate ...
func TestCreate(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error when opening stub DB: %s", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")

	tests := []struct {
		name     string
		data     *models.UserRole
		mockFunc func(dt *models.UserRole)
		wantErr  bool
	}{
		{
			name: "success",
			data: &models.UserRole{
				UserID: uuid.New(),
				RoleID: uuid.New(),
			},
			mockFunc: func(dt *models.UserRole) {
				mocks.ExpectExec("INSERT INTO user_roles (created_at,created_by,created_ip,id,role_id,user_id) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, sqlmock.AnyArg(), dt.RoleID, dt.UserID).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: &models.UserRole{
				UserID: uuid.New(),
				RoleID: uuid.New(),
			},
			mockFunc: func(dt *models.UserRole) {
				mocks.ExpectExec("INSERT INTO user_roles (created_at,created_by,created_ip,id,role_id,user_id) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, sqlmock.AnyArg(), dt.RoleID, dt.UserID).
					WillReturnResult(driver.RowsAffected(1)).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.data)
			r := NewRepository(dbConn)
			err := r.Create(context.Background(), tt.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mocks.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

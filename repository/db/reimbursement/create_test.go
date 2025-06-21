// Package reimbursement ...
package reimbursement

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
		data     *models.Reimbursement
		mockFunc func(dt *models.Reimbursement)
		wantErr  bool
	}{
		{
			name: "success",
			data: &models.Reimbursement{
				EmployeeID:  uuid.New(),
				Amount:      1000,
				Description: "test",
			},
			mockFunc: func(dt *models.Reimbursement) {
				mocks.ExpectExec("INSERT INTO roles (amount,created_at,created_by,created_ip,date,description,employee_id,id) VALUES ($1,NOW(),$2,$3,$4,$5,$6,$7)").
					WithArgs(dt.Amount, dt.CreatedBy, dt.CreatedIP, dt.Date, dt.Description, dt.EmployeeID, sqlmock.AnyArg()).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: &models.Reimbursement{
				EmployeeID:  uuid.New(),
				Amount:      1000,
				Description: "test",
			},
			mockFunc: func(dt *models.Reimbursement) {
				mocks.ExpectExec("INSERT INTO roles (amount,created_at,created_by,created_ip,date,description,employee_id,id) VALUES ($1,NOW(),$2,$3,$4,$5,$6,$7)").
					WithArgs(dt.Amount, dt.CreatedBy, dt.CreatedIP, dt.Date, dt.Description, dt.EmployeeID, sqlmock.AnyArg()).
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

// Package attendance ...
package attendance

import (
	models "attendance/model/db"
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

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
		data     *models.Attendance
		mockFunc func(dt *models.Attendance)
		wantErr  bool
	}{
		{
			name: "success",
			data: &models.Attendance{
				EmployeeID: uuid.New(),
				Date:       time.Now(),
			},
			mockFunc: func(dt *models.Attendance) {
				mocks.ExpectExec("INSERT INTO attendances (created_at,created_by,created_ip,date,employee_id,id) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, dt.Date, dt.EmployeeID, sqlmock.AnyArg()).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: &models.Attendance{
				EmployeeID: uuid.New(),
				Date:       time.Now(),
			},
			mockFunc: func(dt *models.Attendance) {
				mocks.ExpectExec("INSERT INTO attendances (created_at,created_by,created_ip,date,employee_id,id) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, dt.Date, dt.EmployeeID, sqlmock.AnyArg()).
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

// Package attendanceperiod ...
package attendanceperiod

import (
	models "attendance/model/db"
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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
		data     *models.AttendancePeriod
		mockFunc func(dt *models.AttendancePeriod)
		wantErr  bool
	}{
		{
			name: "success",
			data: &models.AttendancePeriod{
				StartDate: time.Now(),
				EndDate:   time.Now().Add(time.Hour * 24 * 25),
			},
			mockFunc: func(dt *models.AttendancePeriod) {
				mocks.ExpectExec("INSERT INTO attendance_periods (created_at,created_by,created_ip,end_date,id,start_date) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, dt.EndDate, sqlmock.AnyArg(), dt.StartDate).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: &models.AttendancePeriod{
				StartDate: time.Now(),
				EndDate:   time.Now().Add(time.Hour * 24 * 25),
			},
			mockFunc: func(dt *models.AttendancePeriod) {
				mocks.ExpectExec("INSERT INTO attendance_periods (created_at,created_by,created_ip,end_date,id,start_date) VALUES (NOW(),$1,$2,$3,$4,$5)").
					WithArgs(dt.CreatedBy, dt.CreatedIP, dt.EndDate, sqlmock.AnyArg(), dt.StartDate).
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

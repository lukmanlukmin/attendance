// Package payroll ...
package payroll

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

// TestUpdate ...
func TestUpdate(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error when opening stub DB: %s", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")

	tests := []struct {
		name     string
		data     *models.Payroll
		mockFunc func(dt *models.Payroll)
		wantErr  bool
	}{
		{
			name: "success",
			data: &models.Payroll{
				ID:                 uuid.New(),
				AttendancePeriodID: uuid.New(),
				Status:             "pending",
			},
			mockFunc: func(dt *models.Payroll) {
				mocks.ExpectExec("UPDATE payrolls SET deleted_at = $1, is_deleted = $2, status = $3, updated_at = NOW() WHERE id = $4").
					WithArgs(dt.DeletedAt, dt.IsDeleted, dt.Status, dt.ID).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: &models.Payroll{
				ID:                 uuid.New(),
				AttendancePeriodID: uuid.New(),
				Status:             "pending",
			},
			mockFunc: func(dt *models.Payroll) {
				mocks.ExpectExec("UPDATE payrolls SET deleted_at = $1, is_deleted = $2, status = $3, updated_at = NOW() WHERE id = $4").
					WithArgs(dt.DeletedAt, dt.IsDeleted, dt.Status, dt.ID).
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
			err := r.Update(context.Background(), tt.data)
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

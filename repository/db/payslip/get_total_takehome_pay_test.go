// Package payslip ...
package payslip

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestGetTotalTakeHomePay ...
func TestGetTotalTakeHomePay(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockPid := uuid.New()
	mockSum := 1000000

	tests := []struct {
		name     string
		pid      uuid.UUID
		mockFunc func(p uuid.UUID)
		wantErr  bool
	}{
		{
			name: "success",
			pid:  mockPid,
			mockFunc: func(p uuid.UUID) {
				rows := sqlmock.NewRows([]string{
					"COALESCE(SUM(p.take_home_pay), 0)",
				}).AddRow(
					mockSum,
				)

				mocks.ExpectQuery(`SELECT COALESCE(SUM(p.take_home_pay), 0) FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3`).
					WithArgs(false, p, false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			pid:  mockPid,
			mockFunc: func(p uuid.UUID) {
				mocks.ExpectQuery(`SELECT COALESCE(SUM(p.take_home_pay), 0) FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3`).
					WithArgs(false, p, false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.pid)
			result, err := r.GetTotalTakeHomePay(context.Background(), tt.pid)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, 0, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, mockSum, result)
			}
		})
	}
}

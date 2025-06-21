// Package payslip ...
package payslip

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

// TestGetByPayrollAndEmployee ...
func TestGetByPayrollAndEmployee(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockID := uuid.New()
	mockEID := uuid.New()
	mockPID := uuid.New()

	tests := []struct {
		name     string
		eID      uuid.UUID
		pID      uuid.UUID
		mockFunc func(eID, pID uuid.UUID)
		wantErr  bool
	}{
		{
			name: "success",
			eID:  mockEID,
			pID:  mockPID,
			mockFunc: func(eID, pID uuid.UUID) {
				rows := sqlmock.NewRows([]string{
					"id", "payroll_id", "employee_id", "base_salary", "attendance_days", "attendance_off_days", "attendance_deduction", "overtime_hours", "overtime_multiply_rate", "overtime_pay", "reimbursement_total", "take_home_pay", "generated_at", "deleted_at", "is_deleted",
				}).AddRow(
					mockID, mockPID, mockEID, 10000, 100, 100, 100, 100, 100, 100, 100, 100, time.Now(), nil, false,
				)

				mocks.ExpectQuery(`SELECT p.id, p.payroll_id, p.employee_id, p.base_salary, p.attendance_days, p.attendance_off_days, p.attendance_deduction, p.overtime_hours, p.overtime_multiply_rate, p.overtime_pay, p.reimbursement_total, p.take_home_pay, p.generated_at, p.deleted_at, p.is_deleted FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id WHERE p.employee_id = $1 AND p.is_deleted = $2 AND pr.id = $3`).
					WithArgs(eID, false, pID).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			eID:  mockEID,
			pID:  mockPID,
			mockFunc: func(eID, pID uuid.UUID) {
				mocks.ExpectQuery(`SELECT p.id, p.payroll_id, p.employee_id, p.base_salary, p.attendance_days, p.attendance_off_days, p.attendance_deduction, p.overtime_hours, p.overtime_multiply_rate, p.overtime_pay, p.reimbursement_total, p.take_home_pay, p.generated_at, p.deleted_at, p.is_deleted FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id WHERE p.employee_id = $1 AND p.is_deleted = $2 AND pr.id = $3`).
					WithArgs(eID, false, pID).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.eID, tt.pID)
			result, err := r.GetByPayrollAndEmployee(context.Background(), tt.pID, tt.eID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, mockID, result.ID)
				assert.Equal(t, mockEID, result.EmployeeID)
				assert.Equal(t, mockPID, result.PayrollID)
			}
		})
	}
}

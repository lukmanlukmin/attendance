// Package payslip ...
package payslip

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

// TestCreateBulk ...
func TestCreateBulk(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error when opening stub DB: %s", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")

	tests := []struct {
		name     string
		data     []models.Payslip
		mockFunc func(dt []models.Payslip)
		wantErr  bool
	}{
		{
			name: "success",
			data: []models.Payslip{
				{
					EmployeeID:           uuid.New(),
					PayrollID:            uuid.New(),
					BaseSalary:           1000000,
					AttendanceDays:       8,
					AttendanceOffDays:    2,
					AttendanceDeduction:  200000,
					OvertimeHours:        0,
					OvertimeMultiplyRate: 1,
					OvertimePay:          0,
					ReimbursementTotal:   0,
					TakeHomePay:          800000,
					GeneratedAt:          time.Now(),
				},
			},
			mockFunc: func(dt []models.Payslip) {
				mocks.ExpectExec("INSERT INTO payslips (id,payroll_id,employee_id,base_salary,attendance_days,attendance_off_days,attendance_deduction,overtime_hours,overtime_multiply_rate,overtime_pay,reimbursement_total,take_home_pay,generated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,NOW())").
					WithArgs(sqlmock.AnyArg(), dt[0].PayrollID, dt[0].EmployeeID, dt[0].BaseSalary, dt[0].AttendanceDays, dt[0].AttendanceOffDays, dt[0].AttendanceDeduction, dt[0].OvertimeHours, dt[0].OvertimeMultiplyRate, dt[0].OvertimePay, dt[0].ReimbursementTotal, dt[0].TakeHomePay).
					WillReturnResult(driver.RowsAffected(1))
			},
			wantErr: false,
		},
		{
			name: "error",
			data: []models.Payslip{
				{
					EmployeeID:           uuid.New(),
					PayrollID:            uuid.New(),
					BaseSalary:           1000000,
					AttendanceDays:       8,
					AttendanceOffDays:    2,
					AttendanceDeduction:  200000,
					OvertimeHours:        0,
					OvertimeMultiplyRate: 1,
					OvertimePay:          0,
					ReimbursementTotal:   0,
					TakeHomePay:          800000,
					GeneratedAt:          time.Now(),
				},
			},
			mockFunc: func(dt []models.Payslip) {
				mocks.ExpectExec("INSERT INTO payslips (id,payroll_id,employee_id,base_salary,attendance_days,attendance_off_days,attendance_deduction,overtime_hours,overtime_multiply_rate,overtime_pay,reimbursement_total,take_home_pay,generated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,NOW())").
					WithArgs(sqlmock.AnyArg(), dt[0].PayrollID, dt[0].EmployeeID, dt[0].BaseSalary, dt[0].AttendanceDays, dt[0].AttendanceOffDays, dt[0].AttendanceDeduction, dt[0].OvertimeHours, dt[0].OvertimeMultiplyRate, dt[0].OvertimePay, dt[0].ReimbursementTotal, dt[0].TakeHomePay).
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
			err := r.CreateBulk(context.Background(), tt.data)
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

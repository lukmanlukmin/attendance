// Package payslip ...
package payslip

import (
	"attendance/model/db"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestGetResumeList ...
func TestGetResumeList(t *testing.T) {
	dbMock, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")

	type args struct {
		payrollID uuid.UUID
		page      uint64
		perPage   uint64
	}
	tests := []struct {
		name       string
		args       args
		mockCount  func(args args)
		mockSelect func(args args)
		wantCount  int
		wantResult []db.ResumePayslip
		expectErr  bool
	}{
		{
			name: "success get resume list",
			args: args{
				payrollID: uuid.New(),
				page:      1,
				perPage:   10,
			},
			mockCount: func(args args) {
				mock.ExpectQuery(`SELECT COUNT(*) FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3`).
					WithArgs(false, args.payrollID, false).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			mockSelect: func(args args) {
				mock.ExpectQuery(`SELECT p.id, p.employee_id, e.full_name, p.take_home_pay FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3 LIMIT 10 OFFSET 0`).
					WithArgs(false, args.payrollID, false).
					WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "full_name", "take_home_pay"}).
						AddRow(uuid.New(), uuid.New(), "John Doe", int64(5000000)))
			},
			wantCount: 1,
			wantResult: []db.ResumePayslip{
				{
					EmployeeFullName: "John Doe",
					TakeHomePay:      5000000,
				},
			},
			expectErr: false,
		},
		{
			name: "error in count query",
			args: args{
				payrollID: uuid.New(),
				page:      1,
				perPage:   10,
			},
			mockCount: func(args args) {
				mock.ExpectQuery(`SELECT COUNT(*) FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3`).
					WithArgs(false, args.payrollID, false).
					WillReturnError(assert.AnError)
			},
			mockSelect: func(_ args) {
			},
			wantCount:  0,
			wantResult: nil,
			expectErr:  true,
		},
		{
			name: "error in select query",
			args: args{
				payrollID: uuid.New(),
				page:      1,
				perPage:   10,
			},
			mockCount: func(args args) {
				mock.ExpectQuery(`SELECT COUNT(*) FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3`).
					WithArgs(false, args.payrollID, false).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			mockSelect: func(args args) {
				mock.ExpectQuery(`SELECT p.id, p.employee_id, e.full_name, p.take_home_pay FROM payslips p JOIN payrolls pr ON pr.id = p.payroll_id JOIN employees e ON e.id = p.employee_id WHERE p.is_deleted = $1 AND pr.id = $2 AND pr.is_deleted = $3 LIMIT 10 OFFSET 0`).
					WithArgs(false, args.payrollID, false).
					WillReturnError(assert.AnError)
			},
			wantCount:  0,
			wantResult: nil,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := NewRepository(dbConn)
			tt.mockCount(tt.args)
			tt.mockSelect(tt.args)

			got, total, err := r.GetResumeList(context.Background(), tt.args.payrollID, tt.args.page, tt.args.perPage)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Zero(t, total)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCount, total)
				assert.Len(t, got, len(tt.wantResult))
				assert.Equal(t, tt.wantResult[0].EmployeeFullName, got[0].EmployeeFullName)
				assert.Equal(t, tt.wantResult[0].TakeHomePay, got[0].TakeHomePay)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

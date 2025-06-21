// Package reimbursement ...
package reimbursement

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestSumEmployeeReimbursements ...
func TestSumEmployeeReimbursements(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockEid := uuid.New()
	mockPid := uuid.New()
	mockSum := 10000

	tests := []struct {
		name     string
		eid      uuid.UUID
		pid      uuid.UUID
		mockFunc func(e, p uuid.UUID)
		wantErr  bool
	}{
		{
			name: "success",
			eid:  mockEid,
			pid:  mockPid,
			mockFunc: func(e, p uuid.UUID) {
				rows := sqlmock.NewRows([]string{
					"COALESCE(SUM(r.amount), 0)",
				}).AddRow(
					mockSum,
				)

				mocks.ExpectQuery(`SELECT COALESCE(SUM(r.amount), 0) FROM reimbursements r JOIN attendance_periods ap ON r.date BETWEEN ap.start_date AND ap.end_date WHERE r.employee_id = $1 AND ap.id = $2 AND r.is_deleted = $3`).
					WithArgs(e, p, false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			eid:  mockEid,
			pid:  mockPid,
			mockFunc: func(e, p uuid.UUID) {
				mocks.ExpectQuery(`SELECT COALESCE(SUM(r.amount), 0) FROM reimbursements r JOIN attendance_periods ap ON r.date BETWEEN ap.start_date AND ap.end_date WHERE r.employee_id = $1 AND ap.id = $2 AND r.is_deleted = $3`).
					WithArgs(e, p, false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.eid, tt.pid)
			result, err := r.SumEmployeeReimbursements(context.Background(), tt.eid, tt.pid)
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

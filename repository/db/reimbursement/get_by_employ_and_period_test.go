// Package reimbursement ...
package reimbursement

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

// TestGetByEmployeeAndPeriod ...
func TestGetByEmployeeAndPeriod(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockEid := uuid.New()
	mockPid := uuid.New()

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
					"id", "employee_id", "amount", "description", "created_at",
				}).AddRow(
					uuid.New(), e, 10000, "description", time.Now(),
				)

				mocks.ExpectQuery(`SELECT r.id, r.employee_id, r.amount, r.description, r.created_at FROM reimbursements r JOIN attendance_periods ap ON ap.id = $1 WHERE ap.is_deleted = $2 AND r.employee_id = $3 AND r.is_deleted = $4 AND r.date BETWEEN ap.start_date AND ap.end_date ORDER BY r.created_at ASC`).
					WithArgs(p, false, e, false).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			eid:  mockEid,
			pid:  mockPid,
			mockFunc: func(e, p uuid.UUID) {
				mocks.ExpectQuery(`SELECT r.id, r.employee_id, r.amount, r.description, r.created_at FROM reimbursements r JOIN attendance_periods ap ON ap.id = $1 WHERE ap.is_deleted = $2 AND r.employee_id = $3 AND r.is_deleted = $4 AND r.date BETWEEN ap.start_date AND ap.end_date ORDER BY r.created_at ASC`).
					WithArgs(p, false, e, false).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.eid, tt.pid)
			result, err := r.GetByEmployeeAndPeriod(context.Background(), tt.eid, tt.pid)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, 0, len(result))
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, 1, len(result))
			}
		})
	}
}

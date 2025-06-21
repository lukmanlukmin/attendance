// Package attendance ...
package attendance

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// TestCountEmployeeAttendance ...
func TestCountEmployeeAttendance(t *testing.T) {
	dbMock, mocks, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()
	dbConn := sqlx.NewDb(dbMock, "sqlmock")
	mockEid := uuid.New()
	mockPid := uuid.New()
	mockSum := 10

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
					"COUNT(1)",
				}).AddRow(
					mockSum,
				)

				mocks.ExpectQuery(`SELECT COUNT(1) FROM attendances a JOIN attendance_periods ap ON a.date BETWEEN ap.start_date AND ap.end_date WHERE a.employee_id = $1 AND a.is_deleted = $2 AND ap.id = $3`).
					WithArgs(e, false, p).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "not found",
			eid:  mockEid,
			pid:  mockPid,
			mockFunc: func(e, p uuid.UUID) {
				mocks.ExpectQuery(`SELECT COUNT(1) FROM attendances a JOIN attendance_periods ap ON a.date BETWEEN ap.start_date AND ap.end_date WHERE a.employee_id = $1 AND a.is_deleted = $2 AND ap.id = $3`).
					WithArgs(e, false, p).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(dbConn)
			tt.mockFunc(tt.eid, tt.pid)
			result, err := r.CountEmployeeAttendance(context.Background(), tt.eid, tt.pid)
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

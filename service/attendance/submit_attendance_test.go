// Package attendance ...
package attendance

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	"attendance/repository/db"
	mockAttendance "attendance/repository/db/attendance/mocks"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestSubmitAttendance ...
func TestSubmitAttendance(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	attendanceMock := mockAttendance.NewMockIAttendance(ctrl)

	tests := []struct {
		name     string
		empID    string
		nowTime  func() time.Time
		mockFunc func()
		errRes   error
		wantErr  bool
	}{
		{
			name:  "success",
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				attendanceMock.EXPECT().IsAttendanceSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				attendanceMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "fail weekend",
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-21 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
			},
			wantErr: true,
			errRes:  constant.ErrWeekendAttendance,
		},
		{
			name:  "fail weekend",
			empID: "",
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
			},
			wantErr: true,
			errRes:  constant.ErrEmployeeNotFound,
		},
		{
			name:  "fail already submit",
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				attendanceMock.EXPECT().IsAttendanceSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
			wantErr: true,
			errRes:  constant.ErrAlreadySubmitAttendance,
		},
		{
			name:  "fail error query 1",
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				attendanceMock.EXPECT().IsAttendanceSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, sql.ErrConnDone)
			},
			wantErr: true,
			errRes:  sql.ErrConnDone,
		},
		{
			name:  "fail error query 2",
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 09:00:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				attendanceMock.EXPECT().IsAttendanceSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				attendanceMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(sql.ErrConnDone)
			},
			wantErr: true,
			errRes:  sql.ErrConnDone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			svc := &Service{
				Repository: &repository.Repository{
					DB: db.Repository{
						Attendance: attendanceMock,
					},
				},
				Config:  &config.Config{},
				NowFunc: tt.nowTime,
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, constant.ContextUserIDKey, uuid.NewString())
			ctx = context.WithValue(ctx, constant.ContextEmployeeIDKey, tt.empID)
			ctx = context.WithValue(ctx, constant.ContextRoleKey, constant.RoleEmployee)
			ctx = context.WithValue(ctx, constant.ContextRequestIDKey, "req-xyz")
			ctx = context.WithValue(ctx, constant.ContextIPKey, "192.168.1.10")

			err := svc.SubmitAttendance(ctx)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Package attendance ...
package attendance

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	"attendance/model/payload"
	"attendance/repository/db"
	mockOvertime "attendance/repository/db/overtime/mocks"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestSubmitOvertime ...
func TestSubmitOvertime(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	overtimeMock := mockOvertime.NewMockIOvertime(ctrl)

	tests := []struct {
		name     string
		data     payload.SubmitOvertimeRequest
		empID    string
		nowTime  func() time.Time
		mockFunc func()
		errRes   error
		wantErr  bool
	}{
		{
			name:  "success",
			empID: uuid.NewString(),
			data:  payload.SubmitOvertimeRequest{Hours: 1},
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:01:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				overtimeMock.EXPECT().IsOvertimeSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				overtimeMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "fail max overtime",
			data:  payload.SubmitOvertimeRequest{Hours: 5},
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:01:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
			},
			wantErr: true,
			errRes:  constant.ErrMaximumOvertime,
		},
		{
			name:  "fail submit before working hour end",
			data:  payload.SubmitOvertimeRequest{Hours: 3},
			empID: uuid.NewString(),
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 16:45:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
			},
			wantErr: true,
			errRes:  constant.ErrSubmitOvertimeBeforeWorkingHour,
		},
		{
			name:  "fail not employee",
			data:  payload.SubmitOvertimeRequest{Hours: 3},
			empID: "",
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:45:00") // nolint:errcheck // skip for test only
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
			data:  payload.SubmitOvertimeRequest{Hours: 1},
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:01:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				overtimeMock.EXPECT().IsOvertimeSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
			wantErr: true,
			errRes:  constant.ErrAlreadySubmitOvertime,
		},
		{
			name:  "error query 1",
			empID: uuid.NewString(),
			data:  payload.SubmitOvertimeRequest{Hours: 1},
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:01:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				overtimeMock.EXPECT().IsOvertimeSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, sql.ErrConnDone)
			},
			wantErr: true,
			errRes:  sql.ErrConnDone,
		},
		{
			name:  "error query 1",
			empID: uuid.NewString(),
			data:  payload.SubmitOvertimeRequest{Hours: 1},
			nowTime: func() time.Time {
				tm, _ := time.Parse("2006-01-02 15:04:05", "2025-06-20 17:01:00") // nolint:errcheck // skip for test only
				return tm
			},
			mockFunc: func() {
				overtimeMock.EXPECT().IsOvertimeSubmitted(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil)
				overtimeMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(sql.ErrConnDone)
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
						Overtime: overtimeMock,
					},
				},
				Config: &config.Config{
					Application: config.Application{
						MaxOvertimeHour: 3,
						EndWorkingHour:  17,
					},
				},
				NowFunc: tt.nowTime,
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, constant.ContextUserIDKey, uuid.NewString())
			ctx = context.WithValue(ctx, constant.ContextEmployeeIDKey, tt.empID)
			ctx = context.WithValue(ctx, constant.ContextRoleKey, constant.RoleEmployee)
			ctx = context.WithValue(ctx, constant.ContextRequestIDKey, "req-xyz")
			ctx = context.WithValue(ctx, constant.ContextIPKey, "192.168.1.10")

			err := svc.SubmitOvertime(ctx, tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

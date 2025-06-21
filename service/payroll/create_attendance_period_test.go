// Package payroll ...
package payroll

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	"attendance/model/payload"
	"attendance/repository/db"
	mockattendancePeriond "attendance/repository/db/attendance_period/mocks"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestCreateAttendancePeriod ...
func TestCreateAttendancePeriod(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	attendancePeriodMock := mockattendancePeriond.NewMockIAttendancePeriod(ctrl)

	tests := []struct {
		name     string
		data     payload.CreateAttendancePeriodRequest
		empID    string
		mockFunc func(data payload.CreateAttendancePeriodRequest)
		errRes   error
		wantErr  bool
	}{
		{
			name:  "success",
			data:  payload.CreateAttendancePeriodRequest{StartDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24 * 25)},
			empID: uuid.NewString(),
			mockFunc: func(data payload.CreateAttendancePeriodRequest) {
				attendancePeriodMock.EXPECT().IsOverLapping(gomock.Any(), data.StartDate, data.EndDate).Return(false, nil)
				attendancePeriodMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.data)
			svc := &Service{
				Repository: &repository.Repository{
					DB: db.Repository{
						AttendancePeriod: attendancePeriodMock,
					},
				},
				Config: &config.Config{
					Application: config.Application{
						AllowOverLapPeriod: false,
					},
				},
				NowFunc: time.Now,
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, constant.ContextUserIDKey, uuid.NewString())
			ctx = context.WithValue(ctx, constant.ContextEmployeeIDKey, tt.empID)
			ctx = context.WithValue(ctx, constant.ContextRoleKey, constant.RoleEmployee)
			ctx = context.WithValue(ctx, constant.ContextRequestIDKey, "req-xyz")
			ctx = context.WithValue(ctx, constant.ContextIPKey, "192.168.1.10")

			err := svc.CreateAttendancePeriod(ctx, tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

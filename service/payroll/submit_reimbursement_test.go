// Package payroll ...
package payroll

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	"attendance/model/payload"
	"attendance/repository/db"
	mockReimbursement "attendance/repository/db/reimbursement/mocks"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestSubmitReimbursement ...
func TestSubmitReimbursement(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reimbursementMock := mockReimbursement.NewMockIReimbursement(ctrl)

	tests := []struct {
		name     string
		data     payload.SubmitReimbursementRequest
		empID    string
		mockFunc func(data payload.SubmitReimbursementRequest)
		errRes   error
		wantErr  bool
	}{
		{
			name:  "success",
			data:  payload.SubmitReimbursementRequest{Amount: 10000, Description: "test"},
			empID: uuid.NewString(),
			mockFunc: func(_ payload.SubmitReimbursementRequest) {
				reimbursementMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "fail not employee",
			data:  payload.SubmitReimbursementRequest{Amount: 10000, Description: "test"},
			empID: "",
			mockFunc: func(_ payload.SubmitReimbursementRequest) {
			},
			wantErr: true,
			errRes:  constant.ErrEmployeeNotFound,
		},
		{
			name:  "error query",
			data:  payload.SubmitReimbursementRequest{Amount: 10000, Description: "test"},
			empID: uuid.NewString(),
			mockFunc: func(_ payload.SubmitReimbursementRequest) {
				reimbursementMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(sql.ErrConnDone)
			},
			wantErr: true,
			errRes:  sql.ErrConnDone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.data)
			svc := &Service{
				Repository: &repository.Repository{
					DB: db.Repository{
						Reimbursement: reimbursementMock,
					},
				},
				Config:  &config.Config{},
				NowFunc: time.Now,
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, constant.ContextUserIDKey, uuid.NewString())
			ctx = context.WithValue(ctx, constant.ContextEmployeeIDKey, tt.empID)
			ctx = context.WithValue(ctx, constant.ContextRoleKey, constant.RoleEmployee)
			ctx = context.WithValue(ctx, constant.ContextRequestIDKey, "req-xyz")
			ctx = context.WithValue(ctx, constant.ContextIPKey, "192.168.1.10")

			err := svc.SubmitReimbursement(ctx, tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

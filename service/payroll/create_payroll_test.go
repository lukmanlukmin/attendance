// Package payroll ...
package payroll

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	modelDB "attendance/model/db"
	"attendance/repository/db"
	mockPayroll "attendance/repository/db/payroll/mocks"
	mockKafkaProducer "attendance/repository/kafka/mocks"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lukmanlukmin/go-lib/database/connection"
	"github.com/stretchr/testify/assert"
)

// TestCreatePayroll ...
func TestCreatePayroll(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	payrollMock := mockPayroll.NewMockIPayroll(ctrl)
	kafkaProducerMock := mockKafkaProducer.NewMockIRepository(ctrl)

	tests := []struct {
		name     string
		pID      uuid.UUID
		mockFunc func(sql sqlmock.Sqlmock, pID uuid.UUID)
		errRes   error
		wantErr  bool
	}{
		{
			name: "success",
			pID:  uuid.New(),
			mockFunc: func(sqlMock sqlmock.Sqlmock, pID uuid.UUID) {
				payrollMock.EXPECT().GetByAttendacePeriod(gomock.Any(), pID, nil).Return([]modelDB.Payroll{}, nil)
				sqlMock.ExpectBegin()
				payrollMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				kafkaProducerMock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				sqlMock.ExpectCommit()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DB, mockedSQL, _ := sqlmock.New()

			defer func() {
				_ = DB.Close()
			}()

			tt.mockFunc(mockedSQL, tt.pID)
			svc := &Service{
				Repository: &repository.Repository{
					DB: db.Repository{
						Payroll: payrollMock,
					},
					KafkaProducer: kafkaProducerMock,
					Store: &connection.Store{
						Master: &sqlx.DB{
							DB: DB,
						},
					},
				},
				Config:  &config.Config{},
				NowFunc: time.Now,
			}

			ctx := context.Background()
			ctx = context.WithValue(ctx, constant.ContextUserIDKey, uuid.NewString())
			ctx = context.WithValue(ctx, constant.ContextRoleKey, constant.RoleEmployee)
			ctx = context.WithValue(ctx, constant.ContextRequestIDKey, "req-xyz")
			ctx = context.WithValue(ctx, constant.ContextIPKey, "192.168.1.10")

			err := svc.CreatePayroll(ctx, tt.pID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

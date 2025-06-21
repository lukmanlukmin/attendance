// Package auth ...
package auth

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/constant"
	modelDB "attendance/model/db"
	"attendance/model/payload"
	"attendance/repository/db"
	mockEmployee "attendance/repository/db/employee/mocks"
	mockRole "attendance/repository/db/role/mocks"
	mockUser "attendance/repository/db/user/mocks"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/util"
	"github.com/stretchr/testify/assert"
)

// TestValidateUserByCredential ...
func TestValidateUserByCredential(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMock := mockUser.NewMockIUser(ctrl)
	roleMock := mockRole.NewMockIRole(ctrl)
	employeeMock := mockEmployee.NewMockIEmployee(ctrl)

	tests := []struct {
		name     string
		data     payload.LoginCredential
		mockFunc func(data payload.LoginCredential)
		errRes   error
		wantErr  bool
	}{
		{
			name: "success",
			data: payload.LoginCredential{Username: "username", Password: "password"},
			mockFunc: func(data payload.LoginCredential) {
				userID := uuid.New()
				passwordHash, _ := util.HashPassword("password")
				mockUser := &modelDB.User{
					ID:       userID,
					Username: "username",
					Password: passwordHash,
				}
				userMock.EXPECT().GetByUsername(gomock.Any(), data.Username).Return(mockUser, nil)

				mockRole := []modelDB.Role{
					{
						Name: constant.RoleAdmin,
					},
				}
				roleMock.EXPECT().GetByUserID(gomock.Any(), userID).Return(mockRole, nil)
				employeeMock.EXPECT().GetByUserID(gomock.Any(), userID).Return(&modelDB.Employee{
					ID:     uuid.New(),
					UserID: uuid.New(),
				}, nil)
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
						User:     userMock,
						Role:     roleMock,
						Employee: employeeMock,
					},
				},
				Conf: &config.Config{
					Security: config.Security{
						JWTSecret:   "secret",
						JWTDuration: "1h",
					},
				},
			}

			ctx := context.Background()

			res, err := svc.ValidateUserByCredential(ctx, tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errRes.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

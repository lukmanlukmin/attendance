// Package auth ...
package auth

import (
	"attendance/constant"
	"attendance/model/payload"
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/lukmanlukmin/go-lib/util"
)

// ValidateUserByCredential ...
func (s *Service) ValidateUserByCredential(ctx context.Context, data payload.LoginCredential) (*payload.Token, error) {
	userData, err := s.Repository.DB.User.GetByUsername(ctx, data.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constant.ErrInvalidUserCredentials
		}
		return nil, err
	}

	if userData != nil && util.CheckPassword(userData.Password, data.Password) {
		jwtData := map[string]interface{}{
			constant.ContextUserIDKey: userData.ID.String(),
		}

		roles, err := s.Repository.DB.Role.GetByUserID(ctx, userData.ID)
		if err != nil {
			return nil, constant.ErrFailedGenerateToken
		}
		roleString := []string{}
		for _, role := range roles {
			roleString = append(roleString, role.Name)
			jwtData[constant.ContextRoleKey] = strings.Join(roleString, ",")
		}

		emp, err := s.Repository.DB.Employee.GetByUserID(ctx, userData.ID)
		if err == nil && emp != nil {
			jwtData[constant.ContextEmployeeIDKey] = emp.ID.String()
		}

		duration, err := time.ParseDuration(s.Conf.Security.JWTDuration)
		if err != nil {
			return nil, constant.ErrFailedGenerateToken
		}
		additionalDuration, err := time.ParseDuration("24h")
		if err != nil {
			return nil, constant.ErrFailedGenerateToken
		}
		jwt, err := util.GenerateJWT(s.Conf.Security.JWTSecret, duration, jwtData)
		if err != nil {
			return nil, constant.ErrFailedGenerateToken
		}
		jwtRefresh, err := util.GenerateJWT(s.Conf.Security.JWTSecret, duration+additionalDuration, jwtData)
		if err != nil {
			return nil, constant.ErrFailedGenerateToken
		}
		return &payload.Token{
			Token:        jwt,
			RefreshToken: jwtRefresh,
		}, nil
	}
	return nil, constant.ErrInvalidUserCredentials
}

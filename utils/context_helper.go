// Package utils ...
package utils

import (
	"attendance/constant"
	"context"

	"github.com/google/uuid"
)

// UserContext ...
type UserContext struct {
	UserID     uuid.UUID
	EmployeeID uuid.NullUUID
	Role       string
	RequestID  string
	IPAddress  string
}

// GetUserContext ...
func GetUserContext(ctx context.Context) UserContext {
	userCtx := UserContext{}

	if val, ok := ctx.Value(constant.ContextUserIDKey).(string); ok {
		if id, err := uuid.Parse(val); err == nil {
			userCtx.UserID = id
		}
	}
	if val, ok := ctx.Value(constant.ContextEmployeeIDKey).(string); ok {
		if id, err := uuid.Parse(val); err == nil {
			userCtx.EmployeeID = uuid.NullUUID{UUID: id, Valid: true}
		}
	}
	if val, ok := ctx.Value(constant.ContextRoleKey).(string); ok {
		userCtx.Role = val
	}
	if val, ok := ctx.Value(constant.ContextRequestIDKey).(string); ok {
		userCtx.RequestID = val
	}
	if val, ok := ctx.Value(constant.ContextIPKey).(string); ok {
		userCtx.IPAddress = val
	}

	return userCtx
}

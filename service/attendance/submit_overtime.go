package attendance

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/model/payload"
	"attendance/utils"
	"context"
	"time"
)

// SubmitOvertime ...
func (s *Service) SubmitOvertime(ctx context.Context, data payload.SubmitOvertimeRequest) error {

	now := s.NowFunc().In(time.Local)
	date := now.Truncate(24 * time.Hour)

	if data.Hours > s.Config.Application.MaxOvertimeHour {
		return constant.ErrMaximumOvertime
	}

	if now.Hour() < s.Config.Application.EndWorkingHour {
		return constant.ErrSubmitOvertimeBeforeWorkingHour
	}

	userCtx := utils.GetUserContext(ctx)
	if !userCtx.EmployeeID.Valid {
		return constant.ErrEmployeeNotFound
	}

	exists, err := s.Repository.DB.Overtime.IsOvertimeSubmitted(ctx, userCtx.EmployeeID.UUID, date)
	if err != nil {
		return err
	}
	if exists {
		return constant.ErrAlreadySubmitOvertime
	}

	return s.Repository.DB.Overtime.Create(ctx, &db.Overtime{
		EmployeeID: userCtx.EmployeeID.UUID,
		Date:       date,
		Hours:      data.Hours,
		CreatedBy:  &userCtx.UserID,
		CreatedIP:  &userCtx.IPAddress,
	})
}

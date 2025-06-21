package attendance

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/utils"
	"context"
	"time"

	"github.com/lukmanlukmin/go-lib/log"
)

// SubmitAttendance ...
func (s *Service) SubmitAttendance(ctx context.Context) error {
	now := s.NowFunc().In(time.Local)
	date := now.Truncate(24 * time.Hour)
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return constant.ErrWeekendAttendance
	}

	userCtx := utils.GetUserContext(ctx)
	if !userCtx.EmployeeID.Valid {
		return constant.ErrEmployeeNotFound
	}

	isSubmitted, err := s.Repository.DB.Attendance.IsAttendanceSubmitted(ctx, userCtx.EmployeeID.UUID, date)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to check attendance submitted")
		return err
	}
	if isSubmitted {
		return constant.ErrAlreadySubmitAttendance
	}

	if err := s.Repository.DB.Attendance.Create(ctx, &db.Attendance{
		Date:       date,
		EmployeeID: userCtx.EmployeeID.UUID,
		CreatedBy:  &userCtx.UserID,
		CreatedIP:  &userCtx.IPAddress,
	}); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create attendance")
		return err
	}

	return nil
}

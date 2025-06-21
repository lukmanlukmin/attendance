// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/model/payload"
	"attendance/utils"
	"context"

	"github.com/lukmanlukmin/go-lib/log"
)

// CreateAttendancePeriod ...
func (s *Service) CreateAttendancePeriod(ctx context.Context, data payload.CreateAttendancePeriodRequest) error {

	if !s.Config.Application.AllowOverLapPeriod {
		isOverlap, err := s.Repository.DB.AttendancePeriod.IsOverLapping(ctx, data.StartDate, data.EndDate)
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("failed to check overlapping")
			return err
		}
		if isOverlap {
			return constant.ErrOverlappingAttendancePeriod
		}
	}

	userCtx := utils.GetUserContext(ctx)
	start, end := utils.NormalizeAttendancePeriod(data.StartDate, data.EndDate)
	ap := &db.AttendancePeriod{
		StartDate: start,
		EndDate:   end,
		CreatedBy: &userCtx.UserID,
		CreatedIP: &userCtx.IPAddress,
	}
	if err := s.Repository.DB.AttendancePeriod.Create(ctx, ap); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create attendance period")
		return err
	}
	return nil
}

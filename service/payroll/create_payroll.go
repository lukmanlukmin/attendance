package payroll

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/model/event"
	"attendance/utils"
	"context"

	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/database"
	"github.com/lukmanlukmin/go-lib/log"
)

// CreatePayroll ...
func (s *Service) CreatePayroll(ctx context.Context, periodID uuid.UUID) error {
	userCtx := utils.GetUserContext(ctx)

	payrolls, err := s.Repository.DB.Payroll.GetByAttendacePeriod(ctx, periodID, nil)
	if err != nil {
		return err
	}
	if len(payrolls) > 0 {
		return constant.ErrPayrollAlreadyRequested
	}

	return database.BeginTransaction(ctx, s.Store.GetMaster(), func(ctx context.Context) error {
		payrollData := &db.Payroll{
			AttendancePeriodID: periodID,
			Status:             constant.StatusPending,
			CreatedBy:          &userCtx.UserID,
			CreatedIP:          &userCtx.IPAddress,
		}
		if err = s.Repository.DB.Payroll.Create(ctx, payrollData); err != nil {
			log.WithContext(ctx).WithError(err).Error("failed to create payroll")
			return err
		}

		return s.KafkaProducer.Publish(ctx, constant.TopicCalculatePayroll, event.CalculatePayrollJob{
			AttendancePeriodID: periodID,
			PayrollID:          payrollData.ID,
		})
	})
}

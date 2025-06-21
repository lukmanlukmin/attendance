package payroll

import (
	"attendance/constant"
	db "attendance/model/db"
	"attendance/model/event"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lukmanlukmin/go-lib/database"
)

// CalculatePayroll ...
func (s *Service) CalculatePayroll(ctx context.Context, data event.CalculatePayrollJob) error {
	batchSize := 50
	offset := 0

	payrols, err := s.Repository.DB.Payroll.GetByAttendacePeriod(ctx, data.AttendancePeriodID, nil)
	if err != nil {
		return err
	}

	if len(payrols) <= 0 {
		return constant.ErrPayrollNotFound
	}

	payrolData := db.Payroll{}
	for _, pr := range payrols {
		if pr.Status != constant.StatusPending {
			pr.Status = constant.StatusFailed
			if err := s.Repository.DB.Payroll.Update(ctx, &pr); err != nil {
				return err
			}
			return constant.ErrPayrollAlreadyProcessed
		}
		if pr.AttendancePeriodID == pr.AttendancePeriodID && pr.ID == data.PayrollID {
			payrolData = pr
			break
		}
	}

	payrolData.Status = constant.StatusInProgress
	if err := s.Repository.DB.Payroll.Update(ctx, &payrolData); err != nil {
		return err
	}

	attPeriod, err := s.Repository.DB.AttendancePeriod.GetByID(ctx, data.AttendancePeriodID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if attPeriod == nil {
		return constant.ErrAttendancePeriodNotFound
	}

	err = database.BeginTransaction(ctx, s.Store.GetMaster(), func(ctx context.Context) error {
		for {
			employees, err := s.Repository.DB.Employee.GetBatch(ctx, batchSize, offset)
			if err != nil {
				return fmt.Errorf("failed to fetch employees: %w", err)
			}
			if len(employees) == 0 {
				break
			}

			var payslips []db.Payslip
			for _, emp := range employees {
				payslip, err := s.GeneratePayslip(ctx, attPeriod, emp.ID, data.PayrollID)
				if err != nil {
					return fmt.Errorf("failed to generate payslip: %w", err)
				}
				payslips = append(payslips, payslip)
			}

			err = s.Repository.DB.Payslip.CreateBulk(ctx, payslips)
			if err != nil {
				return fmt.Errorf("failed to insert payslips: %w", err)
			}

			offset += batchSize
		}
		return nil
	})
	if err != nil {
		payrolData.Status = constant.StatusFailed
		if err := s.Repository.DB.Payroll.Update(ctx, &payrolData); err != nil {
			return err
		}
		return nil
	}

	payrolData.Status = constant.StatusDone
	if err := s.Repository.DB.Payroll.Update(ctx, &payrolData); err != nil {
		return err
	}
	return nil
}

package payroll

import (
	"attendance/model/payload"
	"context"

	"github.com/google/uuid"
)

// GetResumePayslip ...
func (s *Service) GetResumePayslip(ctx context.Context, payrollID uuid.UUID, page, perPage uint64) (*payload.PayslipSummaryResponse, error) {
	resList, count, err := s.Repository.DB.Payslip.GetResumeList(ctx, payrollID, page, perPage)
	if err != nil {
		return nil, err
	}

	sumItem := []payload.PayslipSummaryItem{}
	for _, r := range resList {
		sumItem = append(sumItem, payload.PayslipSummaryItem{
			EmployeeID:   r.EmployeeID,
			EmployeeName: r.EmployeeFullName,
			TakeHomePay:  r.TakeHomePay,
		})
	}

	totalPay, err := s.Repository.DB.Payslip.GetTotalTakeHomePay(ctx, payrollID)
	if err != nil {
		return nil, err
	}

	return &payload.PayslipSummaryResponse{
		Page:          page,
		PerPage:       perPage,
		TotalData:     count,
		TotalTakeHome: totalPay,
		Data:          sumItem,
	}, nil
}

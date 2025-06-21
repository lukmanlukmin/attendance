// Package payroll ...
package payroll

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go -package=mocks

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	db "attendance/model/db"
	"attendance/model/event"
	"attendance/model/payload"
	"context"
	"time"

	"github.com/google/uuid"
)

// IPayroll ...
type IPayroll interface {
	CreateAttendancePeriod(ctx context.Context, data payload.CreateAttendancePeriodRequest) error
	SubmitReimbursement(ctx context.Context, data payload.SubmitReimbursementRequest) error
	CreatePayroll(ctx context.Context, periodID uuid.UUID) error
	CalculatePayroll(ctx context.Context, data event.CalculatePayrollJob) error
	GeneratePayslip(ctx context.Context, period *db.AttendancePeriod, employeeID, payrollID uuid.UUID) (db.Payslip, error)
	GetPayslip(ctx context.Context, payrollID uuid.UUID) (*payload.PayslipDetailResponse, error)
	GetResumePayslip(ctx context.Context, payrollID uuid.UUID, page, perPage uint64) (*payload.PayslipSummaryResponse, error)
}

// Service ...
type Service struct {
	*repository.Repository
	*config.Config
	NowFunc func() time.Time
}

// NewService ...
func NewService(bs *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Repository: bs,
		Config:     cfg,
		NowFunc:    time.Now,
	}
}

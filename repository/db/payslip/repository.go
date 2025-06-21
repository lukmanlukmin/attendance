// Package payslip ...
package payslip

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	model "attendance/model/db"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// IPayslip ...
type IPayslip interface {
	CreateBulk(ctx context.Context, payslips []model.Payslip) error
	GetByPayrollAndEmployee(ctx context.Context, payrollID, employeeID uuid.UUID) (*model.Payslip, error)
	GetResumeList(ctx context.Context, payrollID uuid.UUID, page, perPage uint64) ([]model.ResumePayslip, int, error)
	GetTotalTakeHomePay(ctx context.Context, payrollID uuid.UUID) (int, error)
}

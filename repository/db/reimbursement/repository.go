// Package reimbursement ...
package reimbursement

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

// IReimbursement ...
type IReimbursement interface {
	Create(ctx context.Context, ri *model.Reimbursement) error
	SumEmployeeReimbursements(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) (int, error)
	GetByEmployeeAndPeriod(ctx context.Context, employeeID, attendancePeriodID uuid.UUID) ([]model.Reimbursement, error)
}

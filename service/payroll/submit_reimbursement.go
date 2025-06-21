// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/db"
	"attendance/model/payload"
	"attendance/utils"
	"context"
	"time"
)

// SubmitReimbursement ...
func (s *Service) SubmitReimbursement(ctx context.Context, data payload.SubmitReimbursementRequest) error {

	userCtx := utils.GetUserContext(ctx)
	if !userCtx.EmployeeID.Valid {
		return constant.ErrEmployeeNotFound
	}

	today := s.NowFunc().Truncate(24 * time.Hour)
	return s.Repository.DB.Reimbursement.Create(ctx, &db.Reimbursement{
		EmployeeID:  userCtx.EmployeeID.UUID,
		Date:        today,
		Amount:      data.Amount,
		Description: data.Description,
		CreatedBy:   &userCtx.UserID,
		CreatedIP:   &userCtx.IPAddress,
	})
}

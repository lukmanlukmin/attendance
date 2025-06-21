// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Reimbursement ...
type Reimbursement struct {
	ID          uuid.UUID  `db:"id"`
	EmployeeID  uuid.UUID  `db:"employee_id"`
	Date        time.Time  `db:"date"`
	Amount      int        `db:"amount"`
	Description string     `db:"description"`
	CreatedBy   *uuid.UUID `db:"created_by"`
	CreatedIP   *string    `db:"created_ip"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	IsDeleted   bool       `db:"is_deleted"`
}

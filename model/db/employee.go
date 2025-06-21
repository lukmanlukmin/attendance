// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Employee ...
type Employee struct {
	ID        uuid.UUID  `db:"id"`
	UserID    uuid.UUID  `db:"user_id"`
	FullName  string     `db:"full_name"`
	Salary    int        `db:"salary"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	IsDeleted bool       `db:"is_deleted"`
}

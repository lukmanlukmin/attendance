package db

import (
	"time"

	"github.com/google/uuid"
)

// UserRole ...
type UserRole struct {
	ID        uuid.UUID  `db:"id"`
	UserID    uuid.UUID  `db:"user_id"`
	RoleID    uuid.UUID  `db:"role_id"`
	CreatedBy *uuid.UUID `db:"created_by"`
	UpdatedBy *uuid.UUID `db:"updated_by"`
	CreatedIP *string    `db:"created_ip"`
	UpdatedIP *string    `db:"updated_ip"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	IsDeleted bool       `db:"is_deleted"`
}

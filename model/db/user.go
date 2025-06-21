// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID        uuid.UUID  `db:"id"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	CreatedBy *uuid.UUID `db:"created_by"`
	UpdatedBy *uuid.UUID `db:"updated_by"`
	CreatedIP *string    `db:"created_ip"`
	UpdatedIP *string    `db:"updated_ip"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	IsDeleted bool       `db:"is_deleted"`
}

// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// Role ...
type Role struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

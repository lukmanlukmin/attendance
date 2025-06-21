// Package db ...
package db

import (
	"time"

	"github.com/google/uuid"
)

// AuditLog ...
type AuditLog struct {
	ID        uuid.UUID  `db:"id"`
	TableName string     `db:"table_name"`
	Action    string     `db:"action"` // CREATE, UPDATE, DELETE
	RecordID  uuid.UUID  `db:"record_id"`
	UserID    *uuid.UUID `db:"user_id"`
	IPAddress *string    `db:"ip_address"`
	RequestID *string    `db:"request_id"`
	CreatedAt time.Time  `db:"created_at"`
}

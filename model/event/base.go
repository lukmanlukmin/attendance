// Package event ...
package event

import "time"

// MessageFormat ...
type MessageFormat struct {
	Data     interface{} `json:"data,omitempty"`
	Metadata Metadata    `json:"metadata,omitempty"`
}

// Metadata ...
type Metadata struct {
	EmitHost  string    `json:"emit_host,omitempty"`
	EmitTime  int64     `json:"emit_time,omitempty"`
	Event     string    `json:"event,omitempty"`
	Hash      string    `json:"hash,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

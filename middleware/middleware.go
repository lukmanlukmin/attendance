// Package middleware ...
package middleware

import (
	"attendance/config"
)

// Midleware ...
type Midleware struct {
	Conf *config.Config
}

// New ...
func New(conf *config.Config) *Midleware {
	return &Midleware{
		Conf: conf,
	}
}

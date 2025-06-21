// Package event ...
package event

import "attendance/bootstrap"

// Handler ...
type Handler struct {
	*bootstrap.Bootstrap
}

// NewHandler ...
func NewHandler(bs *bootstrap.Bootstrap) *Handler {
	return &Handler{Bootstrap: bs}
}

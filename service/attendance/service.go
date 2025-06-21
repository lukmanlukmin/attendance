// Package attendance ...
package attendance

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go -package=mocks

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/model/payload"
	"context"
	"time"
)

// IAttendance ...
type IAttendance interface {
	SubmitAttendance(ctx context.Context) error
	SubmitOvertime(ctx context.Context, data payload.SubmitOvertimeRequest) error
}

// Service ...
type Service struct {
	*repository.Repository
	*config.Config
	NowFunc func() time.Time
}

// NewService ...
func NewService(bs *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Repository: bs,
		Config:     cfg,
		NowFunc:    time.Now,
	}
}

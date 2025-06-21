// Package auth ...
package auth

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go -package=mocks

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/model/payload"
	"context"
)

// Service ...
type Service struct {
	*repository.Repository
	Conf *config.Config
}

// NewService ...
func NewService(bs *repository.Repository, conf *config.Config) *Service {
	return &Service{
		Repository: bs,
		Conf:       conf,
	}
}

// IAuth ...
type IAuth interface {
	ValidateUserByCredential(ctx context.Context, data payload.LoginCredential) (*payload.Token, error)
}

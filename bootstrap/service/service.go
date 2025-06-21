// Package service ...
package service

import (
	"attendance/bootstrap/repository"
	"attendance/config"
	"attendance/service"
	"attendance/service/attendance"
	"attendance/service/auth"
	"attendance/service/payroll"

	connDB "github.com/lukmanlukmin/go-lib/database/connection"
)

// Service ...
type Service struct {
	Store *connDB.Store
	service.Service
}

// LoadServices ...
func LoadServices(repo *repository.Repository, conf *config.Config) *Service {
	return &Service{
		Service: service.Service{
			Attendance: attendance.NewService(repo, conf),
			Payroll:    payroll.NewService(repo, conf),
			Auth:       auth.NewService(repo, conf),
		},
	}
}

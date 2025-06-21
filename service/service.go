// Package service ...
package service

import (
	"attendance/service/attendance"
	"attendance/service/auth"
	"attendance/service/payroll"
)

// Service ...
type Service struct {
	Auth       auth.IAuth
	Attendance attendance.IAttendance
	Payroll    payroll.IPayroll
}

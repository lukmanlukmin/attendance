// Package db ...
package db

import (
	"attendance/repository/db/attendance"
	attendanceperiod "attendance/repository/db/attendance_period"
	"attendance/repository/db/employee"
	"attendance/repository/db/overtime"
	"attendance/repository/db/payroll"
	"attendance/repository/db/payslip"
	"attendance/repository/db/reimbursement"
	"attendance/repository/db/role"
	"attendance/repository/db/user"
	userrole "attendance/repository/db/user_role"
)

// Repository ...
type Repository struct {
	User             user.IUser
	Role             role.IRole
	UserRole         userrole.IUserRole
	Employee         employee.IEmployee
	AttendancePeriod attendanceperiod.IAttendancePeriod
	Attendance       attendance.IAttendance
	Overtime         overtime.IOvertime
	Reimbursement    reimbursement.IReimbursement
	Payroll          payroll.IPayroll
	Payslip          payslip.IPayslip
}

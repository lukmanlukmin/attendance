// Package repository ...
package repository

import (
	"attendance/repository/db"
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
	repoKafka "attendance/repository/kafka"

	connDB "github.com/lukmanlukmin/go-lib/database/connection"
	"github.com/lukmanlukmin/go-lib/kafka"
)

// Repository ...
type Repository struct {
	Store         *connDB.Store
	DB            db.Repository
	KafkaProducer repoKafka.IRepository
}

// LoadRepository ...
func LoadRepository(connectionDB *connDB.Store, kafkaProducer kafka.Producer) *Repository {
	return &Repository{
		Store: connectionDB,
		DB: db.Repository{
			User:             user.NewRepository(connectionDB.GetMaster()),
			Role:             role.NewRepository(connectionDB.GetMaster()),
			UserRole:         userrole.NewRepository(connectionDB.GetMaster()),
			Employee:         employee.NewRepository(connectionDB.GetMaster()),
			AttendancePeriod: attendanceperiod.NewRepository(connectionDB.GetMaster()),
			Attendance:       attendance.NewRepository(connectionDB.GetMaster()),
			Overtime:         overtime.NewRepository(connectionDB.GetMaster()),
			Reimbursement:    reimbursement.NewRepository(connectionDB.GetMaster()),
			Payroll:          payroll.NewRepository(connectionDB.GetMaster()),
			Payslip:          payslip.NewRepository(connectionDB.GetMaster()),
		},
		KafkaProducer: repoKafka.NewRepository(kafkaProducer),
	}
}

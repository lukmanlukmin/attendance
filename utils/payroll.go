package utils

import "math"

// CalculateOvertimePay ...
func CalculateOvertimePay(baseSalary int, overtimeHours int, multiplier float64) int {
	const totalWorkingDays = 22
	const workingHoursPerDay = 8

	hourlyRate := float64(baseSalary) / float64(totalWorkingDays*workingHoursPerDay)
	overtimePay := float64(overtimeHours) * hourlyRate * multiplier

	return int(math.Round(overtimePay))
}

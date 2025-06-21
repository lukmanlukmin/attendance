// Package utils ...
package utils

import "time"

// NormalizeAttendancePeriod ...
func NormalizeAttendancePeriod(start, end time.Time) (time.Time, time.Time) {
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), end.Location())
	return start, end
}

// CountWeekdays ...
func CountWeekdays(start, end time.Time) int {
	start = start.Truncate(24 * time.Hour)
	end = end.Truncate(24 * time.Hour)

	if end.Before(start) {
		start, end = end, start
	}

	count := 0
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			count++
		}
	}
	return count
}

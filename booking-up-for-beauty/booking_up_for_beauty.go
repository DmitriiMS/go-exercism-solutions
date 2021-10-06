package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// 7/25/2019 13:45:00
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	today := time.Now()
	appointment, _ := time.Parse(layout, date)
	return appointment.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	noon := rounded.Add(time.Hour * 12)
	evening := rounded.Add(time.Hour * 18)
	return t.After(noon) && t.Before(evening)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYearOfOurLordAndSaviorJesusChrist := time.Now().Year()
	aniversary := time.Date(currentYearOfOurLordAndSaviorJesusChrist, 9, 15, 0, 0, 0, 0, time.UTC)
	return aniversary
}

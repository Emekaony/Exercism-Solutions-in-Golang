package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// the layout must look just like the string coming in
	// down to the number of 0 paddings and whether is is 24-hr or 12-hr
	tt, _ := time.Parse("1/02/2006 15:04:05", date)
	return tt
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	givenTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(givenTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	// something that finally worked for the first time!
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	message := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return message
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

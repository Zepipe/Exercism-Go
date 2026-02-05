package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    timeParsed, _ := time.Parse("1/2/2006 15:04:05", date)
    return timeParsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    timeParsed, _ := time.Parse("January 2, 2006 15:04:05", date)
    return time.Now().After(timeParsed)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	timeParsed, _ := time.Parse("Monday, March 2, 2006 15:04:05", date)
    return timeParsed.Hour() >= 12 && timeParsed.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {

    timeParsed, _ := time.Parse("1/2/2006 15:04:05", date)
    
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
        timeParsed.Weekday(),
        timeParsed.Month(),
        timeParsed.Day(),
        timeParsed.Year(),
        timeParsed.Hour(),
        timeParsed.Minute(),
    )
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    now := time.Now()
    return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

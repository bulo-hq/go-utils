package iso8601

import "time"

// Parse parses an ISO 8601 formatted string into a time.Time object.
func Parse(s string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", s)
}

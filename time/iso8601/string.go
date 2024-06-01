package iso8601

import "time"

// String returns a time.Time object as an ISO 8601 formatted string.
func String(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-07:00")
}

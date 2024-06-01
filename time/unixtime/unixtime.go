package unixtime

import "time"

// UnixTime returns a time.Time object as a Unix timestamp.
func UnixTime(t *time.Time) *int64 {
	if t == nil {
		return nil
	}
	ut := t.Unix()
	return &ut
}

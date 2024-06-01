package unixtime

import "time"

// Time returns a Unix timestamp as a time.Time object.
func Time(ut *int64) *time.Time {
	if ut == nil {
		return nil
	}
	t := time.Unix(*ut, 0)
	return &t
}

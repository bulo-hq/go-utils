package iso8601

import "time"

func Time(datetimeString string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", datetimeString)
}

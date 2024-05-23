package iso8601

import "time"

func String(datetime time.Time) string {
	return datetime.Format("2006-01-02T15:04:05-07:00")
}

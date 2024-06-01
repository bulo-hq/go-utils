package gtfstime

import "time"

// Parse parses a start date and start time into a time.Time object.
func Parse(startDate string, startTime string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", startDate+" "+startTime)
	if err != nil {
		return nil, err
	}
	tlocal := t.Local()
	return &tlocal, nil
}

package domain

import (
	"errors"
	"math"
	"time"
)

// Period is an immutable time-span object representing
// the start and end dates of type time.Time
// Not comparable with ==.
type Period struct {
	start time.Time
	end   time.Time
}

// ErrInvalidPeriod error occurs when the entered data is invalid;
// the end is before the start, or the end is the same as the start.
var ErrInvalidPeriod = errors.New("period: the entered data is invalid")

// NewPeriod validates input data;
// checks entries for the Err Invalid Period error.
func NewPeriod(start, end time.Time) (Period, error) {
	if end.Before(start) || start.Equal(end) {
		return Period{}, ErrInvalidPeriod
	}
	return Period{start: start, end: end}, nil
}

// Days calculates the number of days
// between the start and end of the Period structure;
// Rounds the number of days up
func (p Period) Days() int64 {
	duration := p.end.Sub(p.start)
	days := math.Ceil(duration.Hours() / 24)
	return int64(days)
}

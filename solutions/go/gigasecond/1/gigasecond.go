// Package gigasecond adds a gigasecond time unit.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the time one gigasecond after its given time parameter.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}

package snapshot

import (
	"time"
)

// UTCTime returns a Time object in UTC
func UTCTime(timestamp uint64) time.Time {
	return time.Unix(int64(timestamp), 0).UTC()
}

// SameDateOfTimestamps checks a timestamp for date equality
func SameDateOfTimestamps(ts1, ts2 uint64) bool {
	date1 := UTCTime(ts1).Format("02-01-2006")
	date2 := UTCTime(ts2).Format("02-01-2006")
	return (date1 == date2)
}

// AddressUnion returns the union of two Ethereum address maps
func AddressUnion(a, b map[string]bool) map[string]bool {
	for k := range b {
		a[k] = true
	}
	return a
}

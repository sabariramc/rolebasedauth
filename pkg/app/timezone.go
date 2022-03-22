package app

import "time"

var tz *time.Location

func GetTimeZone() *time.Location {
	return tz
}

func setTimeZone(t *time.Location) {
	tz = t
}

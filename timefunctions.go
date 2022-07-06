package swordtech

import (
	"fmt"
	"time"
)

// stringToTime attempts to convert the string to a valid time. Nil is returned if this is unsuccessful
func stringToTime(v string) *time.Time {
	var yyyy, mm, dd, hh, mn, ss int
	n, err := fmt.Sscanf(v, "%4d-%2d-%2dT%2d:%2d:%2d", &yyyy, &mm, &dd, &hh, &mn, &ss)
	if n != 6 || err != nil {
		n, err = fmt.Sscanf(v, "%4d-%2d-%2d %2d:%2d:%2d", &yyyy, &mm, &dd, &hh, &mn, &ss)
		if n != 6 || err != nil {
			return nil
		}
	}
	t := time.Date(yyyy, time.Month(mm), dd, hh, mn, ss, 0, time.Local)
	return &t
}

// TtmeToString converts a time value to a SQL string representation
func timeToString(v *time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute(), v.Second())
}

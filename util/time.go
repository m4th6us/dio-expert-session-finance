package util

import "time"

func StringToTime(value string) time.Time {

	var layout = "2006-01-02T15:04:05"
	convertedTime, _ := time.Parse(layout, value)

	return convertedTime

}

package datetime

import (
	"time"
)

// DateToStr converts Date object to String
func DateToStr(layout string, date time.Time) string {

	if layout == "" {
		layout = "2006-01-02"
	}
	return date.Format(layout)
}

// EpochToStr converts unix epoch time  to String
func EpochToStr(layout string, epochSec *int64) string {

	if layout == "" {
		layout = "2006-01-02"
	}
	return time.Unix(*epochSec, 0).Format(layout)
}

// TimeToStr converts time object to String
func TimeToStr(layout string, time time.Time) string {

	if layout == "" {
		layout = "2006-01-02T15:04:05.000Z"
	}
	return time.Format(layout)
}

// StrToTime converts string  to time
func StrToTime(layout string, timeStr string) time.Time {

	if layout == "" {
		layout = "2006-01-02T15:04:05.000Z"
	}
	resp, _ := time.Parse(layout, timeStr)

	return resp
}

package time

import (
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func SecondsFormat(seconds int64) string {
	return SecondsToTime(seconds).Format(TimeFormat)
}

func MillisecondsFormat(milliseconds int64) string {
	return MillisecondsToTime(milliseconds).Format(TimeFormat)
}

func SecondsToTime(seconds int64) time.Time {
	return MillisecondsToTime(seconds * 1000)
}

func MillisecondsToTime(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*1e6)
}

func GetTodayStartTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

package utils

import (
	"github.com/BitofferHub/xtimer/internal/constant"
	"time"
)

func GetStartMinute(timeStr string) (time.Time, error) {
	return time.ParseInLocation(constant.MinuteFormat, timeStr, time.Local)
}

func GetDayStr(t time.Time) string {
	return t.Format(constant.DayFormat)
}

func GetHourStr(t time.Time) string {
	return t.Format(constant.HourFormat)
}

func GetMinuteStr(t time.Time) string {
	return t.Format(constant.MinuteFormat)
}

func GetStartHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
}

func GetMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

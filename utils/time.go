package utils

import "time"

func GetTodayDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
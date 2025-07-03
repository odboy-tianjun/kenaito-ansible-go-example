package util

import "time"

const layoutDateTime = "2006-01-02 15:04:05"

func FormatDateTime(time time.Time) string {
	return time.Format(layoutDateTime)
}

package utils

import "time"

// FormatDate 时间格式化
func FormatDate(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

// ParseDate 解析时间
func ParseDate(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
}

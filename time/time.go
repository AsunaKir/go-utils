package main

import (
	"fmt"
	"time"
)

// GetCurrentDateTime 获取当前时间，格式为：2006-01-02 15:04:05
func GetCurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTime 获取当前时间，格式为：2006-01-02
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02")
}

// GetCurrentTimestamp 获取当前时间戳，秒级 eg:1758192515
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// GetCurrentTimestampMilli 获取当前时间戳，毫秒级 eg:1758192515157
func GetCurrentTimestampMilli() int64 {
	return time.Now().UnixMilli()
}

// TimestampToDateTime 时间转换:时间戳 --> YYYY-MM-DD HH:mm:ss
func TimestampToDateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// TimestampToDate 时间转换：时间戳 --> YYYY-MM-DD
func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}

// DateTimeToTimestamp 时间转换：YYYY-MM-DD HH:mm:ss --> 时间戳
func DateTimeToTimestamp(dateTime string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// FormatYYYYMMDD 时间格式转换:YYYYMMDD -> YYYY-MM-DD eg: 20060102 --> 2006-01-02
func FormatYYYYMMDD(timeStr string) (string, error) {
	t, err := time.Parse("20060102", timeStr)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02"), nil
}

// FormatYYYYMMDDReverse 时间格式转换:YYYY-MM-DD -> YYYYMMDD eg: 2006-01-02 --> 20060102
func FormatYYYYMMDDReverse(timeStr string) (string, error) {
	t, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		return "", err
	}
	return t.Format("20060102"), nil
}

// GetFirstAndLastDayOfMonth 获取指定YYYY-MM的第一天和最后一天 eg: 2025-09 --> 2025-09-01,2025-09-30
func GetFirstAndLastDayOfMonth(yearMonth string) (string, string, error) {
	t, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		return "", "", err
	}
	return t.AddDate(0, 0, 0).Format("2006-01-02"), t.AddDate(0, 1, -1).Format("2006-01-02"), nil
}

// GetFirstAndLastDateTimeOfMonth 获取指定YYYY-MM的第一天0时0分0秒和最后一天23时59分59秒的YYYY-MM-DD HH:mm:ss
func GetFirstAndLastDateTimeOfMonth(yearMonth string) (string, string, error) {
	t, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		return "", "", err
	}
	firstTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastTime := firstTime.AddDate(0, 1, 0).Add(-1 * time.Second)
	return firstTime.Format("2006-01-02 15:04:05"), lastTime.Format("2006-01-02 15:04:05"), nil
}

// GetFirstAndLastTimestampOfMonth 获取指定YYYY-MM的第一天0时0分0秒和最后一天23时59分59秒的时间戳
func GetFirstAndLastTimestampOfMonth(yearMonth string) (int64, int64, error) {
	t, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		return 0, 0, err
	}
	firstTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastTime := firstTime.AddDate(0, 1, 0).Add(-1 * time.Second)
	return firstTime.Unix(), lastTime.Unix(), nil
}

// GetStartAndEndDateTimeByDate 获取指定YYYY-MM-DD的0时0分0秒和23时59分59秒的YYYY-MM-DD HH:mm:ss eg: 2025-09-01 --> 2025-09-01 00:00:00,2025-09-01 23:59:59
func GetStartAndEndDateTimeByDate(date string) (string, string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", "", err
	}
	startTime := t.Format("2006-01-02 15:04:05")
	endTime := t.Add(24 * time.Hour).Add(-1 * time.Second).Format("2006-01-02 15:04:05")
	return startTime, endTime, nil
}

// GetStartAndEndTimestampByDate 获取指定YYYY-MM-DD的0时0分0秒和23时59分59秒时间戳
func GetStartAndEndTimestampByDate(date string) (int64, int64, error) {
	t, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return 0, 0, err
	}
	startTime := t.Unix()
	endTime := t.Add(24 * time.Hour).Add(-1 * time.Second).Unix()
	return startTime, endTime, nil
}

func main() {
	fmt.Println(GetStartAndEndTimestampByDate("2025-09-01"))
}

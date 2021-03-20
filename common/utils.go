package common

import (
	"time"
)

type JSONTime time.Time

const (
	// DateFormat        = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

func TotalPage(size, total int64) int64 {
	if size == 0 {
		return 0
	}
	t := total / size
	if total%size > 0 {
		t += 1
	}
	return t
}

func TransPayloadDatetimeFieldValue(value string) (time.Time, error) {
	if value == "" {
		value = DateTimeFormat
	}
	result, err := time.ParseInLocation(DateTimeFormat, value, time.Local)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}

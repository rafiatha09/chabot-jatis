package util

import (
	"time"
)


func GenerateCurrentTimestamp()int64 {
	return time.Now().Unix()
}

func ValidateTimestamp(timestampValidated int64) bool {
	currentTime := GenerateCurrentTimestamp()
	return currentTime > timestampValidated;
}

func IsFutureTime(dateTime time.Time) bool {
    now := time.Now()
    return dateTime.After(now)
}
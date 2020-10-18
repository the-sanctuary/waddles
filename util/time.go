package util

import "time"

var startTime *time.Time

func Uptime() *time.Duration {
	since := time.Since(*startTime)
	return &since
}

func MarkStartTime() {
	now := time.Now()
	startTime = &now
}

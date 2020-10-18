package util

import "time"

var startTime *time.Time

// Uptime returns the time since the bot started
func Uptime() *time.Duration {
	since := time.Since(*startTime)
	return &since
}

// MarkStartTime stores the start time of the bot
func MarkStartTime() {
	now := time.Now()
	startTime = &now
}

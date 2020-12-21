package util

import "time"

var startTime time.Time

// Uptime returns the time since the bot started
func Uptime() time.Duration {
	return time.Since(startTime)
}

// MarkStartTime stores the start time of the bot
func MarkStartTime() {
	startTime = time.Now()
}

package util

import (
	"math"
	"os"
	"os/signal"
	"syscall"
)

//RegisterTermSignals  -
func RegisterTermSignals() {
	// Register term signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

//SliceContains returns true if str is included in slice
func SliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

//AbsInt returns the absolute value of an integer
func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}

// FileExists checks if a path exists and is a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

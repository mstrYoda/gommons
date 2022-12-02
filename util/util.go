package util

import "time"

func ElapsedTime(f func()) time.Duration {
	startTime := time.Now()
	f()
	return time.Now().Sub(startTime)
}

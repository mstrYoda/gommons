package time

import "time"

func ElapsedTime(f func()) time.Duration {
	startTime := time.Now()
	f()
	return time.Now().Sub(startTime)
}

func GivenTimeIsBetweenTimeRange(givenTime, time1, time2 time.Time) bool {
	if givenTime.After(time1) && givenTime.Before(time2) {
		return true
	}
	return false
}

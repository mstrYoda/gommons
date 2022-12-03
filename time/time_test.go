package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestElapsedTime(t *testing.T) {
	elapsedTime := ElapsedTime(func() {
		time.Sleep(100 * time.Millisecond)
	})

	assert.LessOrEqual(t, time.Duration(100), elapsedTime)
}

func TestIsTimeBetweenRange_TimeIsBetweenTimeRange(t *testing.T) {
	givenTime := time.Date(2022, 11, 30, 11, 32, 12, 0, time.Local)
	time1 := time.Date(2022, 11, 30, 10, 31, 12, 0, time.Local)
	time2 := time.Date(2022, 11, 30, 12, 31, 12, 0, time.Local)

	assert.Equal(t, true, IsTimeBetweenRange(givenTime, time1, time2))
}

func TestIsTimeBetweenRange_TimeIsNotBetweenTimeRange(t *testing.T) {
	givenTime := time.Date(2022, 11, 30, 13, 32, 12, 0, time.Local)
	time1 := time.Date(2022, 11, 30, 10, 31, 12, 0, time.Local)
	time2 := time.Date(2022, 11, 30, 12, 31, 12, 0, time.Local)

	assert.Equal(t, false, IsTimeBetweenRange(givenTime, time1, time2))
}

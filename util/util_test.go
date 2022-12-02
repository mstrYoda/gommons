package util

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

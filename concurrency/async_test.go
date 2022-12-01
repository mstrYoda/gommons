package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask(t *testing.T) {
	var a = 0
	var b = 0

	New().Task(
		func() {
			a = 1
			fmt.Println("1")
		}, func() {
			b = 1
			fmt.Println("2")
		}).Await()

	assert.Equal(t, 1, b)
	assert.Equal(t, 1, a)
}

func TestTaskWithResult(t *testing.T) {
	results := NewAsyncWorkWithResult[int]().TaskWithResult(
		func() int {
			return 5
		}, func() int {
			return 11
		}).AwaitResult()

	assert.Contains(t, results, 5)
	assert.Contains(t, results, 11)
}

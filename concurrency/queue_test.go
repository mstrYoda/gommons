package concurrency

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()

	var returnData []int

	q.Publish(context.Background(), 1)

	done := make(chan struct{})
	q.Subscribe(context.Background(), func(data int) {
		fmt.Println("data readed ", data)
		returnData = append(returnData, data)
		done <- struct{}{}
	})
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	q.Publish(ctx, 2)

	assert.Len(t, returnData, 1)
	assert.Eventually(t, func() bool {
		return int32(1) == q.TimeOutMessageCount()
	}, 10*time.Second, 10*time.Millisecond)
}

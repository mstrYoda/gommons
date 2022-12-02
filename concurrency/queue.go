package concurrency

import (
	"context"
	"sync/atomic"
)

type Queue[T any] struct {
	ch                  chan T
	timeOutMessageCount atomic.Int32
	done                chan struct{}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{ch: make(chan T)}
}

func (q *Queue[T]) Publish(ctx context.Context, data T) {
	go func() {
		select {
		case q.ch <- data:
			break
		case <-ctx.Done():
			q.timeOutMessageCount.Add(1)
			break
		}
	}()
}

func (q *Queue[T]) Subscribe(ctx context.Context, fn func(data T)) {
	go func() {
		select {
		case readedData := <-q.ch:
			fn(readedData)
		case <-ctx.Done():
			break
		}
	}()
}

func (q *Queue[T]) TimeOutMessageCount() int32 {
	return q.timeOutMessageCount.Load()
}

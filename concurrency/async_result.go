package concurrency

import (
	"sync"
)

type asyncWorkWithResult[T any] struct {
	tasks []func() T
	wg    sync.WaitGroup
}

func NewAsyncWorkWithResult[T any]() *asyncWorkWithResult[T] {
	return &asyncWorkWithResult[T]{}
}

func (a *asyncWorkWithResult[T]) TaskWithResult(fn ...func() T) *asyncWorkWithResult[T] {
	a.wg.Add(len(fn))
	a.tasks = append(a.tasks, fn...)
	return a
}

func (a *asyncWorkWithResult[T]) AwaitResult() []T {
	results := make([]T, len(a.tasks))

	for i, t := range a.tasks {
		go func(index int, taskFn func() T) {
			defer a.wg.Done()
			results[index] = taskFn()
		}(i, t)
	}

	a.wg.Wait()
	return results
}

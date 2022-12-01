package concurrency

import (
	"sync"
)

type asyncWork struct {
	tasks []func()
	wg    sync.WaitGroup
}

func New() *asyncWork {
	work := &asyncWork{}
	return work
}

func (a *asyncWork) Task(fn ...func()) *asyncWork {
	a.wg.Add(len(fn))
	a.tasks = append(a.tasks, fn...)
	return a
}

func (a *asyncWork) Await() {
	for _, t := range a.tasks {
		go func(taskFn func()) {
			defer a.wg.Done()
			taskFn()
		}(t)
	}

	a.wg.Wait()
}

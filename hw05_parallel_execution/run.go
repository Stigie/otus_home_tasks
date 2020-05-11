package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	// Place your code here
	wg := &sync.WaitGroup{}

	tasksChan := make(chan Task)

	var errorCounter int32

	for i := 0; i < N; i++ {
		wg.Add(1) //nolint:gomnd
		go func(taskChan <-chan Task, wg *sync.WaitGroup) {
			defer wg.Done()
			for task := range taskChan {
				if err := task(); err != nil {
					atomic.AddInt32(&errorCounter, 1)
				}
			}
		}(tasksChan, wg)
	}

	taskCounter := 0

	for taskCounter < len(tasks) && (atomic.LoadInt32(&errorCounter) < int32(M) || M <= 0) {
		tasksChan <- tasks[taskCounter]
		taskCounter++
	}
	close(tasksChan)

	wg.Wait()

	if errorCounter >= int32(M) && M > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

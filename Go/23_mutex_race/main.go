package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func main() {
	counter := Counter{}
	wg := &sync.WaitGroup{}

	// With goroutines - without locks
	// ==================
	// WARNING: DATA RACE
	// exit status 66
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			counter.Increment()
			wg.Done()
		}(wg)
	}

	wg.Wait()

	fmt.Println("Without Locks - Final Counter Value:", counter.count)

	// Resetting the counter for goroutine demo
	counter.count = 0
	mutex := &sync.Mutex{}

	// With goroutines - with locks
	// No Race Condition here.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
			mutex.Lock()
			counter.Increment()
			mutex.Unlock()
			wg.Done()
		}(wg, mutex)
	}

	wg.Wait()

	fmt.Println("With Locks - Final Counter Value:", counter.count)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int)
	myWg := &sync.WaitGroup{}

	myWg.Add(2)

	// Goroutine 1 - Send data to the channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 3; i++ {
			fmt.Println("Sending", i)
			ch <- i
		}
	}(myCh, myWg)

	// Goroutine 2 - Receive data from the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			val, ok := <-ch
			if !ok {
				fmt.Println("Channel Closed")
				return
			}
			fmt.Println("Received", val)
		}
	}(myCh, myWg)

	myWg.Wait()
}

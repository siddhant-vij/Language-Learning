package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var websites = []string{
	"https://google.com",
	"https://facebook.com",
	"https://youtube.com",
	"https://github.com",
	"https://stackoverflow.com",
	"https://amazon.com",
	"https://wikipedia.org",
}

var wg sync.WaitGroup

func getStatusCode(url string, goroutine bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting %s: %v\n", url, err)
		return
	}

	if resp.StatusCode != -1 {
		fmt.Printf("Website: %s, Status Code: %d\n", url, resp.StatusCode)
	}

	if goroutine {
		defer wg.Done()
	}
}

func main() {
	start := time.Now()
	for _, url := range websites {
		getStatusCode(url, false)
	}
	fmt.Printf("\n\nTime taken 1:%v\n\n", time.Since(start))
	// Time taken 1: 4.7803741s

	start = time.Now()
	for _, url := range websites {
		wg.Add(1)
		go getStatusCode(url, true)
	}
	wg.Wait()
	fmt.Printf("\n\nTime taken 2:%v\n\n", time.Since(start))
	// Time taken 2: 616.2825ms
}

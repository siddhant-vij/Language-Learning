package main

import (
	"fmt"
	"time"
)

func main() {
	// Getting the current date and time
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	// Formatting dates and times
	fmt.Println("Formatted date and time:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Formatted date:", currentTime.Format("2006-01-02"))
	fmt.Println("Formatted time:", currentTime.Format("15:04:05"))
	fmt.Println("Day of the week:", currentTime.Format("Monday"))

	// Creating a date and time from scratch
	specificTime := time.Date(2023, time.March, 14, 21, 34, 58, 0, time.Now().Location())
	fmt.Println("Specific time:", specificTime)

	// Using constants for formatting
	fmt.Println("RFC1123Z formatted:", currentTime.Format(time.RFC1123Z))
}

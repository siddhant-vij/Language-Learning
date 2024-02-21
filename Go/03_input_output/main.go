package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Taking input for the user's name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Removing the newline character

	// Taking input for the user's age
	fmt.Print("Enter your age: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr) // Converting string to int
	if err != nil {
		fmt.Println("Invalid input for age. Please enter a number.")
		return
	}
	if age <= 0 {
		fmt.Println("Invalid input for age. Please enter a positive number.")
		return
	}

	// Output the entered information
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

package main

import (
	"fmt"
)

// Employee is a struct to demonstrate methods
type Employee struct {
	Name     string
	Position string
}

// UpdatePosition is a method that updates the Position field of an Employee
func (e *Employee) UpdatePosition(newPosition string) {
	e.Position = newPosition
}

// Sum is a function that takes variable number of arguments and returns their sum
func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// If else demo
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Switch Case demo
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Loops demo
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip the rest of this iteration
		}
		if i == 4 {
			break // Exit the loop
		}
		fmt.Println("Loop iteration:", i)
	}

	// Break, continue and goto demo
	i := 0
Loop:
	for {
		if i == 2 {
			i++
			continue
		}
		if i == 5 {
			break
		}
		if i == 3 {
			i++
			goto Loop
		}
		fmt.Println("Loop with goto iteration:", i)
		i++
	}

	// Functions & Methods demo
	alice := Employee{"Alice", "Developer"}
	fmt.Println("Initial position:", alice.Position)

	alice.UpdatePosition("Senior Developer")
	fmt.Println("Updated position:", alice.Position)

	// Variable Args as parameters demo
	fmt.Println("Sum of numbers:", Sum(1, 2, 3, 4, 5))
	fmt.Println("Sum of numbers:", Sum(1, 2, 3))

	// Anonymous functions demo
	func(message string) {
		fmt.Println("Anonymous function says:", message)
	}("Hello, World!")

	// First Order Functions demo (functions as values and arguments)
	multiply := func(x int, y int) int {
		return x * y
	}
	fmt.Println("Product:", multiply(5, 6))

	// Defer statements demo
	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")
}

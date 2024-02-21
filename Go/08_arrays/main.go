package main

import "fmt"

func main() {
	// Declare an array of 5 integers, initialized with some values
	var numbers = [5]int{10, 20, 30, 40, 50}

	// Print the entire array
	fmt.Println("Numbers:", numbers)

	// Access individual elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])

	// Iterate over the array using a for loop
	fmt.Println("Iterating over the array:")
	for i, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	// Modify an element of the array
	numbers[2] = 35
	fmt.Println("Updated numbers:", numbers)

	// Declare a multi-dimensional array
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi-dimensional array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

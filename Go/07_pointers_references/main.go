package main

import "fmt"

func swapByValue(a int, b int) {
	fmt.Println("Before swapping:", a, b)
	a, b = b, a
	fmt.Println("After swapping:", a, b)
}

func swapByPointers(a *int, b *int) {
	fmt.Println("Before swapping:", *a, *b)
	*a, *b = *b, *a
	fmt.Println("After swapping:", *a, *b)
}

func main() {
	// Declare an integer variable
	var a int = 42

	// Create a pointer that refers to the memory address of variable 'a'
	var b *int = &a

	// Print the value of 'a' and the memory address where 'a' is stored
	fmt.Println("Value of a:", a)    // Output: 42
	fmt.Println("Address of a:", &a) // Output: Some memory address, e.g., 0xc0000120e8

	// Print the memory address stored in pointer 'b' and the value at that memory address (dereferencing)
	fmt.Println("Address stored in pointer b:", b) // Output: Same as the address of 'a'
	fmt.Println("Value pointed to by b:", *b)      // Output: 42

	// Changing the value at the memory address pointed to by 'b' (dereferencing and assignment)
	*b = *b + 2
	fmt.Println("New value of a after change through pointer b:", a) // Output: 44

	// ---------------------------------------------

	// Declare two integer variables
	var x int = 1
	var y int = 2

	// Swap by value
	swapByValue(x, y)
	fmt.Printf("After swap by value: x = %d, y = %d\n", x, y)

	// Swap by pointers
	swapByPointers(&x, &y)
	fmt.Printf("After swap by pointers: x = %d, y = %d\n", x, y)
}

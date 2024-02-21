package main

import (
	"fmt"
	"sort"
)

func main() {
	// Using the make function to create a slice with length 5 and capacity 10
	sliceFromMake := make([]int, 5, 10)
	fmt.Println("Slice from make with length 5 and capacity 10:", sliceFromMake, "Length:", len(sliceFromMake), "Capacity:", cap(sliceFromMake))

	// The new function is not generally used with slices since it returns a pointer to the zero value of the type.
	// For slices, the zero value is nil, so using new with slices is not common - but can be done:
	sliceFromNew := new([]int)
	fmt.Println("Slice from new:", sliceFromNew, "Length:", len(*sliceFromNew), "Capacity:", cap(*sliceFromNew))

	// Declaration and usage of a slice
	var mySlice []int                  // Declaration without initialization
	mySlice = append(mySlice, 1, 2, 3) // Appending values to the slice
	fmt.Println("Declared and initialized slice:", mySlice)

	// Length and capacity of a slice
	fmt.Println("Length of mySlice:", len(mySlice), "Capacity of mySlice:", cap(mySlice))

	// Creating a slice from an array
	array := [5]int{0, 1, 2, 3, 4}
	sliceFromArray := array[1:4] // Creating a slice that includes elements at index 1, 2, and 3
	fmt.Println("Slice from array:", sliceFromArray)

	// Access elements of a slice
	fmt.Println("First element of sliceFromArray:", sliceFromArray[0])

	// Updating the slices
	sliceFromArray[0] = 10 // Changing the value at index 0 of the slice
	fmt.Println("Updated sliceFromArray:", sliceFromArray)

	// Adding values
	sliceFromArray = append(sliceFromArray, 5) // Appending a single value to the slice
	fmt.Println("Slice with an added value:", sliceFromArray)

	// Append one slice to another slice
	anotherSlice := []int{6, 7, 8}
	combinedSlice := append(sliceFromArray, anotherSlice...) // Appending all elements of anotherSlice to sliceFromArray
	fmt.Println("Combined slice:", combinedSlice)

	// Removing value based on index (example: remove the element at index 2)
	indexToRemove := 2
	removedSlice := append(combinedSlice[:indexToRemove], combinedSlice[indexToRemove+1:]...)
	fmt.Println("Slice after removing index 2:", removedSlice)

	// Sorting a slice in increasing order
	sortSlice := []int{5, 2, 8, 1, 9}
	sort.Ints(sortSlice)
	fmt.Println("Sorted slice:", sortSlice)

	// Sorting a slice in decreasing order
	// With Slice() - sorting could be unstable
	// Use SliceStable() for stable sorting instead
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i] > sortSlice[j]
	})
	fmt.Println("Sorted slice in decreasing order:", sortSlice)
}

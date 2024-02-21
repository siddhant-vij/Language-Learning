package main

import (
	"fmt"
)

func main() {
	// Creation of a map
	// employeeSalaries := make(map[string]float64)

	// Creation & initialization of a map with values
	employeeSalaries := map[string]float64{
		"Alice": 60000.00,
		"Bob":   55000.00,
		"Carol": 75000.00,
	}

	// Accessing elements
	bobSalary := employeeSalaries["Bob"]
	fmt.Println("Bob's salary:", bobSalary)

	// Checking if a key exists
	charlieSalary, ok := employeeSalaries["Charlie"]
	if !ok {
		fmt.Println("Charlie's salary not found.")
	} else {
		fmt.Println("Charlie's salary:", charlieSalary)
	}

	// Adding elements to a map
	employeeSalaries["Dan"] = 80000.00

	// Looping through the map entries
	for name, salary := range employeeSalaries {
		fmt.Printf("%s earns %.2f annually\n", name, salary)
	}

	// Deleting an element
	delete(employeeSalaries, "Alice")

	// Updating an element
	employeeSalaries["Bob"] = 58000.00

	// Using the len function to get the number of key-value pairs
	fmt.Println("Total number of employees:", len(employeeSalaries))

	// Maps are reference types
	newMap := employeeSalaries
	newMap["Carol"] = 78000.00 // This will change the original map as well
	fmt.Println("Updated Carol's salary in the original map:", employeeSalaries["Carol"])

	// Creating a nested map
	departments := map[string]map[string]float64{
		"Engineering": {
			"Alice": 120000.00,
			"Bob":   110000.00,
		},
		"Marketing": {
			"Carol": 115000.00,
			"Dan":   130000.00,
		},
	}
	// Accessing nested maps
	fmt.Println("Engineering department salaries:", departments["Engineering"])
}

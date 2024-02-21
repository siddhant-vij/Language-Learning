package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct named 'Employee'
type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

func main() {
	// Instantiation of the struct
	alice := Employee{
		Name:     "Alice",
		Age:      30,
		Position: "Software Developer",
		Salary:   75000.00,
	}

	// Usage of a struct by accessing data members
	fmt.Printf("Employee name: %s\n", alice.Name)
	fmt.Printf("Employee age: %d\n", alice.Age)

	// Pretty printing using fmt package
	fmt.Printf("Employee details: %+v\n", alice)

	// Pretty printing using encoding/json package for a JSON representation
	jsonBytes, err := json.MarshalIndent(alice, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("Employee JSON:\n", string(jsonBytes))

	// Other operations with structs
	// Updating struct field
	alice.Position = "Senior Software Developer"
	fmt.Printf("Updated employee position: %s\n", alice.Position)

	// Instantiation with a pointer to a struct using the built-in 'new' function
	bob := new(Employee)
	bob.Name = "Bob"
	bob.Age = 28
	fmt.Printf("New employee using pointer: %+v\n", *bob)

	// Instantiation with a pointer to a struct using the address-of operator &
	carol := &Employee{"Carol", 32, "Product Manager", 85000.00}
	fmt.Printf("Another employee using pointer: %+v\n", *carol)

	// Anonymous structs
	tempEmployee := struct {
		ID        int
		FirstName string
		LastName  string
	}{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Printf("Temporary employee details: %+v\n", tempEmployee)
}

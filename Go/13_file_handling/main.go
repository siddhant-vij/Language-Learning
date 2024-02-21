package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Employee is a simple struct with relevant data members
type Employee struct {
	Name     string
	Position string
}

// checkNilErr is a helper function to handle errors gracefully
func checkNilErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		os.Exit(1)
	}
}

// writeFile writes the given content to a file with the specified filename
func writeFile(filename string, content string) {
	file, err := os.Create(filename) // Create a file or truncate if it already exists
	checkNilErr(err)
	defer file.Close()

	_, err = file.WriteString(content) // Write the content to the file
	checkNilErr(err)
}

// appendFile appends the given content to a file with the specified filename
func appendFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) // Open the file with write and append mode
	checkNilErr(err)
	defer file.Close()

	_, err = file.WriteString(content) // Append the content to the file
	checkNilErr(err)
}

// readFile reads the content from the specified file and returns it
func readFile(filename string) string {
	file, err := os.Open(filename) // Open the file for reading
	checkNilErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	var content string
	for {
		line, err := reader.ReadString('\n') // Read the file line by line
		if err == io.EOF {
			break
		}
		checkNilErr(err)
		content += line
	}
	return content
}

// readFullFile reads the entire content from the specified file and returns it
func readFullFile(filename string) string {
	content, err := os.ReadFile(filename) // Read the entire file
	checkNilErr(err)
	return string(content)
}

func main() {
	// Define an Employee
	employee := Employee{Name: "John Doe", Position: "Software Engineer"}

	// Convert Employee to string for file operations
	employeeData := fmt.Sprintf("Name: %s\nPosition: %s\n", employee.Name, employee.Position)

	// Write data to a new file
	writeFile("employee.txt", employeeData)

	// Append additional data to the file
	appendFile("employee.txt", "Appended information\n")

	// Read and print the content of the file
	fmt.Println("Reading file line by line:")
	fmt.Println(readFile("employee.txt"))

	// Read the entire file content at once and print it
	fmt.Println("\nReading the whole file at once:")
	fmt.Println(readFullFile("employee.txt"))
}

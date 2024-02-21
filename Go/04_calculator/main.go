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
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1Str = strings.TrimSpace(num1Str)
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		fmt.Println("Invalid input for the first number.")
		return
	}

	fmt.Print("Enter the operator (+, -, *, /, %): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		fmt.Println("Invalid input for the second number.")
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}
}

func calculate(num1 int, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2, nil
	case "%":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 % num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

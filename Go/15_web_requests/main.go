package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// const getUrl = "https://httpbin.org/get"
// const postUrl = "https://httpbin.org/post"
const getUrl = "http://localhost:3000/get"
const postUrl = "http://localhost:3000/post"

func getRequest() {
	// Send GET request
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response from", getUrl, ":\n", string(body))
}

func postRequest() {
	// Prepare the data to post
	data := strings.NewReader(`{
        "UserName": "John Doe",
        "Email": "johndoe@example.com",
        "PhoneNo": "1234567890",
        "PizzaSize": "Large",
        "PizzaToppings": ["Pepperoni", "Mushrooms", "Extra Cheese"],
        "PreferredDeliveryTime": "7:00 PM",
        "DeliveryInstructions": "Ring the bell twice."
    }`)

	// Send POST request
	req, err := http.NewRequest("POST", postUrl, data)
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response from", postUrl, ":\n", string(body))
}

func main() {
	getRequest()
	postRequest()
}

package main

import (
	"fmt"

	"go-learning-19/consume"
	"go-learning-19/create"
)

func testingCreateJson() {
	// Sample Course struct
	course1 := create.Course{
		Username: "john_doe",
		Price:    99.99,
		Platform: "Coursera",
		Password: "strong-password",
		Tags:     []string{"programming", "golang", "json"},
	}

	// Encode the Course struct to JSON
	jsonData1, err := create.EncodeJson(course1)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Prettify the JSON with lowercase keys
	prettyJSON1, err := create.PrettifyJSON(jsonData1)
	if err != nil {
		fmt.Println("Error prettifying JSON:", err)
		return
	}

	// Print the prettified JSON
	fmt.Println(string(prettyJSON1))

	// Trying course2 for omitempty []tags
	course2 := create.Course{
		Username: "john_doe",
		Price:    99.99,
		Platform: "Coursera",
		Password: "strong-password",
		Tags:     nil,
	}

	// Encode the Course struct to JSON
	jsonData2, err := create.EncodeJson(course2)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Prettify the JSON with lowercase keys
	prettyJSON2, err := create.PrettifyJSON(jsonData2)
	if err != nil {
		fmt.Println("Error prettifying JSON:", err)
		return
	}

	// Print the prettified JSON
	fmt.Println(string(prettyJSON2))
}

func testingConsumeJson() {
	// Sample JSON data for Course struct
	sampleJSON := []byte(`{
		"username": "john_doe",
		"price": 99.99,
		"platform": "Coursera",
		"tags": ["programming", "golang", "json"]
	}`)

	// Decode the sample JSON data using DecodeJson function
	decodedCourse, err := consume.DecodeJson(sampleJSON)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the decoded Course object
	fmt.Printf("Decoded Course:\n%+v\n", decodedCourse)
}

func main() {
	testingCreateJson()
	testingConsumeJson()
}

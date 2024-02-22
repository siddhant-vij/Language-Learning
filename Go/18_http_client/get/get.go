package get

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequest(getUrl string) {
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

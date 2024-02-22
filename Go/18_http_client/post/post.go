package post

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func PostRequest(postUrl string, jsonData []byte) {
    // Prepare the data to post
    data := bytes.NewBuffer(jsonData)

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

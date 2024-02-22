package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func deconstructURLWithFormData(encodedURL string) (string, map[string]interface{}) {
	u, err := url.Parse(encodedURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", nil
	}

	formData := make(map[string]interface{})
	for key, values := range u.Query() {
		formData[key] = values[0]
	}

	u.RawQuery = ""

	return u.String(), formData
}

func FormRequest(encodedFormDataUrl string) {
	// Get formData and URL without query parameters
	formDataUrl, formData := deconstructURLWithFormData(encodedFormDataUrl)

	// Convert form data to JSON
	jsonData, err := json.Marshal(formData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send FORM request with JSON data
	req, err := http.NewRequest("FORM", formDataUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating FORM request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making FORM request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response from", formDataUrl, ":\n", string(body))
}

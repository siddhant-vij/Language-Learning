package create

import (
	"encoding/json"
)

// Course struct representing a course
type Course struct {
	Username string   `json:"username"`
	Price    float64  `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson(course Course) ([]byte, error) {
	jsonData, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func PrettifyJSON(jsonData []byte) ([]byte, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	// Convert keys to lowercase
	lowercaseData := make(map[string]interface{})
	for key, value := range data {
		lowercaseData[key] = value
		delete(data, key)
	}

	prettyJSON, err := json.MarshalIndent(lowercaseData, "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJSON, nil
}

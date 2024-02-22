package consume

import "encoding/json"

// Define the Course struct
type Course struct {
	Username string   `json:"username"`
	Price    float64  `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func DecodeJson(jsonData []byte) (*Course, error) {
	course := &Course{}

	err := json.Unmarshal(jsonData, course)
	if err != nil {
		return nil, err
	}

	checkValid := json.Valid(jsonData)
	if checkValid {
		return course, nil
	} else {
		return nil, err
	}
}

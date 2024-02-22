package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	responseData := struct {
		Message string `json:"message"`
	}{
		Message: "Welcome to the Go HTTP server",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jData, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}

	w.Write(jData)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postData map[string]interface{}
	err := decoder.Decode(&postData)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	responseData := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "Received POST data successfully",
		Data:    postData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jData, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}

	w.Write(jData)
}

func handleFormRequest(w http.ResponseWriter, r *http.Request) {
	var formData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	responseData := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "Received Form data successfully",
		Data:    formData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jData, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}

	w.Write(jData)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleGetRequest(w, r)
		case "POST":
			handlePostRequest(w, r)
		case "FORM":
			handleFormRequest(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on localhost:3000...")
	http.ListenAndServe(":3000", nil)
}

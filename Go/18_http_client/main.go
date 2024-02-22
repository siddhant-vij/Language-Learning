package main

import (
	"flag"
	"fmt"
	"net/http"

	"go-learning-18/form"
	"go-learning-18/get"
	"go-learning-18/post"
)

var jsonData = []byte(`{
    "UserName": "ABCD",
    "Email": "abcd@example.com",
    "PhoneNo": "9876543210"
}`)

var baseURL = "http://localhost:3000"
var getUrl = baseURL + "/get"
var postUrl = baseURL + "/post"
var formUrl = baseURL + "/form?name=John&age=30"

func main() {
	getFlag := flag.Bool("get", false, "Perform a GET request")
	postFlag := flag.Bool("post", false, "Perform a POST request")
	formFlag := flag.Bool("form", false, "Perform a FORM request")

	flag.Parse()

	if *getFlag {
		get.GetRequest(getUrl)
	} else if *postFlag {
		post.PostRequest(postUrl, jsonData)
	} else if *formFlag {
		form.FormRequest(formUrl)
	} else {
		fmt.Println("Please provide a flag: -get, -post, or -form")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}

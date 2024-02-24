package main

import (
	"fmt"
	"net/http"

	"go-learning-21/router"
)

func main() {
	r := router.Router()
	fmt.Println("Listening on port 4000")
	http.ListenAndServe(":4000", r)
}

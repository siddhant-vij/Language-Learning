package main

import (
	"fmt"
	"net/url"
	"strings"
)

func constructURL() {
	// Constructing a URL from parts
	scheme := "https"
	host := "example.com"
	path := "/path/to/resource"
	query := ""
	queryParams := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for key, value := range queryParams {
		query += fmt.Sprintf("%s=%s&", key, value)
	}
	query = strings.TrimSuffix(query, "&")
	fmt.Println("Constructed Query:", query)
	urlStr := fmt.Sprintf("%s://%s%s?%s", scheme, host, path, query)

	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Constructed URL:", u)
}

func deconstructURL() {
	// Deconstructing a URL into various parts
	rawURL := "https://example.com/path/to/resource?key1=value1&key2=value2"

	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.RawQuery)
	for k, v := range u.Query() {
		fmt.Printf("%s: %s\n", k, v[0])
	}
}

func main() {
	// Constructing a URL from parts
	fmt.Println("Constructing a URL from parts:")
	constructURL()

	// Deconstructing a URL into various parts
	fmt.Println("\nDeconstructing a URL into various parts:")
	deconstructURL()
}

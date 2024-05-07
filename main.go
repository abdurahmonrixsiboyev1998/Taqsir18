package main

import (
	"fmt"
	"net/url"
)

func ParseURL(inputURL string) (*url.URL, error) {
	num, error := url.Parse(inputURL)
	if error != nil {
		return nil, error
	}
	return num, nil
}

func URLComponents(a *url.URL) {
	fmt.Println("Scheme:", a.Scheme)
	fmt.Println("Host:", a.Host)
	fmt.Println("Path:", a.Path)
	fmt.Println("Query params:")
	Params := a.Query()
	for key, values := range Params {
		fmt.Printf("\t%s: %v\n", key, values)
	}
	fmt.Println("Fragment:", a.Fragment)
}

func QueryParams(a *url.URL) map[string][]string {
	return a.Query()
}

func isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}

func main() {
	URL1 := "https://example.com/path/to/resource?name=Abdurahmon&age=26#section1"
	fmt.Println("Tahlil qilish URL:", URL1)
	URL, err := ParseURL(URL1)
	if err != nil {
		fmt.Println("Tahlil qilish Error URL:", err)
		return
	}
	URLComponents(URL)

	fmt.Println()

	fmt.Println("So'rov parametrlari:")
	query := QueryParams(URL)
	for key, values := range query {
		fmt.Printf("\t%s: %v\n", key, values)
	}

	fmt.Println()

	URL3 := "https://example.com"
	fmt.Printf("Bu '%s' URL to'g'rimi? %t\n", URL3, isValidURL(URL3))

	testInvalidURL := "notavalidurl"
	fmt.Printf("BU '%s' URL to'g'rimi? %t\n", testInvalidURL, isValidURL(testInvalidURL))
}

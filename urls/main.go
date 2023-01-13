package main

import (
	"fmt"
	"net/url"
)

const urls = "https://localhost:3000/test?query=search"

func main() {
	fmt.Println("hello world")
	fmt.Println(urls)

	// parsing
	result, _ := url.Parse(urls)

	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type qparams:%T", qparams)
}

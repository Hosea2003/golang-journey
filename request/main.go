package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nRespone type: %T\n", response)

	defer response.Body.Close() //Must close mannualy

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}

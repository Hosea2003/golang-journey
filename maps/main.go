package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println(languages)

	// delete an item by its key
	delete(languages, "Rb")
	fmt.Println(languages)

	// loops through map
	for _, value := range languages {
		fmt.Printf("value: %v\n", value)
	}
}

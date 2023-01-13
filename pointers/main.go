package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	var number int = 2
	var ptr = &number

	// fmt.Println("value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value of pointer", number)

}

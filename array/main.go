package main

import "fmt"

func main() {
	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Ananas"

	fmt.Println("Fruits", fruits)

	var vegetables = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegetables", vegetables)
}

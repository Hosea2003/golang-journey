package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"Mango", "Banana"}
	fmt.Printf("Type of fruits %T\n", fruits)

	fruits = append(fruits, "Coco", "Apple")
	fmt.Print(fruits)

	fruits = append(fruits[1:])
	fmt.Print("Finally", fruits)

	highScores := make([]int, 4)

	highScores[0] = 1
	highScores[1] = 3
	highScores[2] = 4
	highScores[3] = 2

	sort.Ints(highScores)
	fmt.Println("Highscore", highScores[len(highScores)-1])

	// remove an item by its index
	var courses = []string{"C#", "PHP", "Javascript", "Java"}
	fmt.Println(courses)

	var index int = 3

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

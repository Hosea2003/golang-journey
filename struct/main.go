package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	rindra := User{"rindra", "rindra@gmail.com", true, 20}
	fmt.Println(rindra)
	rindra.setName("ranto")
	rindra.name = "rindra"
	rindra.printName()
}

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) printName() {
	fmt.Println("My name is:", user.name)
}

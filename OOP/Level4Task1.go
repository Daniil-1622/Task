package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name,
		age,
	}
}

func main() {
	User := Person{"Jack", 15}
	NewUser := NewPerson(User.name, User.age)
	fmt.Println(NewUser)
}

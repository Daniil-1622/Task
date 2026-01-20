package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm <%s>", p.Name)
}

func main() {
	User := Person{"Jack", 20}
	result := User.Greet()

	fmt.Println("Приветсвие:", result)
}

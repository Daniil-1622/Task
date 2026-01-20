package main

import "fmt"

type Person struct {
	Name string
	age  int
}

//func (p Person) Greet() string {
//	return fmt.Sprintf("Hello, I'm <%s>", p.Name)
//}

func (p Person) BirthdayValue() {
	p.age++
}
func (p *Person) BirthdayPointer() {
	p.age++
}
func (p Person) GetAge() int {
	return p.age
}

func main() {
	User := Person{"Jack", 20}
	fmt.Println("До изменений:", User.GetAge())

	User.BirthdayValue()
	fmt.Println("Меняем возраст по значению:", User.GetAge()) // Потому что по значению мы работаем с копией и значения теряются

	User.BirthdayPointer()
	fmt.Println("Меняем возраст по указателю:", User.GetAge()) //Потому что работаем с одним участком памяти,возраст увеличивается
}

package main

import "fmt"

type Animal struct{}

func (a Animal) Speak() { fmt.Printf("Звук!") }

type Dog struct {
	Animal
}

func (d Dog) Speak() { fmt.Printf("Woof!") }

type Cat struct {
	Animal
}

func (c Cat) Speak() { fmt.Println("Moow!") }

func main() {
	Dog := Dog{}
	Dog.Speak()
	Cat := Cat{}
	Cat.Speak()
}

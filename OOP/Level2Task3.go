package main

import "fmt"

type One struct{}

func (o One) Print() { fmt.Printf("1") }

type Two struct{}

func (t Two) Print() { fmt.Printf("2") }

type Three struct {
	One
	Two
}

func (t Three) Print() { t.One.Print() }
func main() {
	Chisl := Three{}
	Chisl.Print()
}

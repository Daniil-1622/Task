package main

import "fmt"

type da1 interface {
	Shape()
}

type Square struct {
	a int
	b int
}

func (s Square) Shape() {
	fmt.Println(s.a * s.b)
}

func check(i any) {
	sq, ok := i.(Square)
	if ok {
		fmt.Println("Да")
		sq.Shape()
	} else {
		fmt.Println("Не понял")
	}
}

func main() {
	var s da1 = Square{2, 10}
	check(s)
}

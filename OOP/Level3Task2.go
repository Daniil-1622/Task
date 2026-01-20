package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	long  float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.long * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	restangle := Rectangle{10, 5}
	PrintArea(restangle)
	circle := Circle{10}
	PrintArea(circle)
}

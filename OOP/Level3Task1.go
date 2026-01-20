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

func main() {
	restangle := Rectangle{10, 5}
	fmt.Println("Площадь Rectangle:", restangle.Area())
	circle := Circle{10}
	fmt.Println("Радиус Circle:", circle.Area())
}

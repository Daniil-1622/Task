package main

import "fmt"

type Engine struct{}

func (e Engine) Start() { fmt.Printf("Start") }

type Car struct {
	Engine
}

func (e Car) Drive() { e.Engine.Start() }

func main() {
	car := Car{}
	car.Drive()
}

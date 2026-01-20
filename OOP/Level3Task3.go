package main

import "fmt"

func Describe(v ...interface{}) {
	for _, item := range v {
		switch t := item.(type) {
		case int:
			fmt.Println("Это Int", t)
		case string:
			fmt.Println("Это String", t)
		default:
			fmt.Println("Неизвестно", t)
		}
	}
}

func main() {
	Describe(10, 10.5, true)
	slice := []any{10, "строка"}
	Describe(slice)
}

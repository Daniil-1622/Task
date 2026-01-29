package main

import "fmt"

func main() {
	// Создаем небуферизованный канал типа int
	ch := make(chan int)

	// Отправляем в его число 42
	go func() {
		ch <- 42
	}()

	// Получаем наше число
	value := <-ch

	// Выводим его
	fmt.Println("Received:", value)
}

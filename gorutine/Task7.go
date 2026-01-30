package main

import "fmt"

func main() {
	// Создаем буферизованный канал на 2 значения типа string
	chanel := make(chan string, 2)

	go func() {
		// Записываем сначала "first"
		chanel <- "first"
		// Затем записываем "second"
		chanel <- "second"
	}()
	// Читаем первый элемент
	ch1 := <-chanel
	// Затем второй элемент
	ch2 := <-chanel

	// Тоесть теперь я могу их вывести по очереди
	// Сначала "first"
	fmt.Println("Первый", ch1)
	// Затем "second"
	fmt.Println("Второй", ch2)
}

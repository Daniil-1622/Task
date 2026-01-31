package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем sync.WaitGroup{}
	wg := sync.WaitGroup{}
	// Тут создаем буфферезированный канал
	ch := make(chan int, 3)
	// Добавляем 2 горутины
	wg.Add(2)
	go func() {
		// Откладываем выполнения
		defer wg.Done()
		// Записываем в канала числа: 1, 2, 3
		ch <- 1
		ch <- 2
		ch <- 3
		// Закрываем канал
		close(ch)
	}()

	go func() {
		// Откладываем выполнения
		defer wg.Done()
		// Проходимся циклом по каналу ch
		for v := range ch {
			// Выводим числа
			fmt.Println(v)
		}
	}()
	wg.Wait()
}

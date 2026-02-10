package main

import (
	"fmt"
	"sync"
)

// Создаем функцию для обьединения двух каналов в один
func Merge(ch1, ch2 <-chan int) <-chan int {
	// Используем синхронизацию WaitGroup{}
	wg := sync.WaitGroup{}
	// Создаем канал для результата
	out := make(chan int)
	// Добавляем две задачи нашей WaitGroup{}
	wg.Add(2)

	// Запускаем первую gorutine для чтения из первого канала
	go func() {
		go func() {
			// Откладываем выполнения задачи для WaitGroup{}
			defer wg.Done()
			// Читаем из нашего канала с помощью range
			for n := range ch1 {
				// Записываем его в канал для результата
				out <- n
			}
		}()
		// Чтения из второго канала
		go func() {
			defer wg.Done()
			for n := range ch2 {
				out <- n
			}
		}()
		// Ожидаем выполения WaitGroup{}
		wg.Wait()
		// Закрываем наш общий канал
		close(out)
	}()
	return out
}

func main() {
	// Создаем два канала
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Используем нашу функцию Merge
	merge := Merge(ch1, ch2)

	// Отправка в первый канал
	go func() {
		ch1 <- 55
		close(ch1)
	}()
	// Отправка во второй канал
	go func() {
		ch2 <- 44
		close(ch2)
	}()

	// Проходимся по merge и выводим все значения наших двух каналов
	for val := range merge {
		fmt.Println(val)
	}
}
